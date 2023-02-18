package logx

import (
	"context"
	"gin-admin-api/pkg/pathx"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type (
	Logx struct {
		*logrus.Logger
		conf Conf
	}
	Hook logrus.Hook
)
type Logger = *logrus.Entry

const (
	LogFileMode               = "file"
	LogConsoleMode            = "console"
	LogBothMode               = "both"
	LogFileRotationDaily      = "daily"
	LogFileRotationLevel      = "level"
	LogFileRotationDailyLevel = "level-daily"
	LogFileRotationSize       = "size"
	LogFileJsonEncoding       = "json"
)

var (
	logOnce sync.Once
	log     *Logx

	levelMap = map[string]logrus.Level{
		"panic": logrus.PanicLevel,
		"fatal": logrus.FatalLevel,
		"warn":  logrus.WarnLevel,
		"info":  logrus.InfoLevel,
		"debug": logrus.DebugLevel,
		"trace": logrus.TraceLevel,
	}
)

func Setup(c Conf, opts ...HookOption) {
	if c.Mode == "" {
		log.Fatal("log mode is not set")
	}
	if c.Mode == LogFileMode || c.Mode == LogBothMode {
		logPath := c.Path
		if logPath == "" {
			log.Fatal("log path is not set")
		}
		if ok := pathx.FileExist(logPath); !ok {
			if err := os.MkdirAll(logPath, 0775); err != nil {
				log.Fatalf("log path create fail: %s, %s", logPath, err.Error())
			}
		}
		opts = append([]HookOption{WithGlobalFieldHook, WithLfsHook}, opts...)
	}

	level, ok := levelMap[strings.ToLower(c.Level)]
	if !ok {
		level = logrus.DebugLevel
	}
	logOnce.Do(func() {
		logger := logrus.New()
		log = &Logx{logger, c}
		log.SetLevel(level)

		if c.Mode == LogConsoleMode || c.Mode == LogBothMode {
			log.SetOutput(os.Stdout)
		}
		for _, opt := range opts {
			log.AddHook(opt())
		}
	})
}

func WithContext(ctx context.Context) Logger {
	return log.WithContext(ctx)
}

func WithField(key string, value interface{}) Logger {
	return log.WithField(key, value)
}

func WithFields(fields logrus.Fields) Logger {
	return log.WithFields(fields)
}

func AddGlobalField(key string, value interface{}) {
	globalFields.mut.Lock()
	defer globalFields.mut.Unlock()
	globalFields.fields[key] = value
}

func AddGlobalFields(fields logrus.Fields) {
	globalFields.mut.Lock()
	defer globalFields.mut.Unlock()
	for k, v := range fields {
		globalFields.fields[k] = v
	}
}

func AddHook(hook Hook) {
	log.AddHook(hook)
}

func (l *Logx) getFormatter() logrus.Formatter {
	var formatter logrus.Formatter
	timeFormat := "2006-01-02 15:04:05.999"
	if l.conf.Encoding == LogFileJsonEncoding {
		formatter = &logrus.JSONFormatter{
			TimestampFormat: timeFormat,
		}
	} else {
		formatter = &logrus.TextFormatter{
			TimestampFormat: timeFormat,
		}
	}
	return formatter
}

func (l *Logx) fileRotate(levelVal string) *rotatelogs.RotateLogs {
	filePath := l.conf.FilePrefix
	if l.conf.Rotation == LogFileRotationLevel || l.conf.Rotation == LogFileRotationDailyLevel {
		filePath += "-" + levelVal
	}
	if l.conf.Rotation == LogFileRotationDaily || l.conf.Rotation == LogFileRotationDailyLevel {
		filePath += "-%Y-%m-%d"
	}
	filePath = strings.TrimLeft(filePath, "-") + l.conf.FileSuffix
	fileFullPath := path.Join(l.conf.Path, filePath)
	options := make([]rotatelogs.Option, 0)

	if l.conf.KeepDays > 0 {
		// 文件最大保存时间
		options = append(options, rotatelogs.WithMaxAge(time.Duration(l.conf.KeepDays)*time.Hour*24))
	}
	if l.conf.RotationTime > 0 {
		// 日志切割时间间隔
		options = append(options, rotatelogs.WithRotationTime(time.Duration(l.conf.KeepDays)*time.Hour))
	}
	if l.conf.Rotation == LogFileRotationSize && l.conf.MaxSize > 0 {
		options = append(options, rotatelogs.WithRotationSize(int64(l.conf.MaxSize)*1024*1024))
	}
	// 软连接
	// fileRotateLogs.WithLinkName(logPath)
	writer, err := rotatelogs.New(fileFullPath, options...)
	if err != nil {
		log.Fatalf("log file rotate Error: %s", err.Error())
	}
	return writer
}
