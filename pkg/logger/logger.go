package logger

import (
	"gin-admin-api/pkg/utils"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type LogConf struct {
	Path             string // 日志存储地址
	Level            string // 日志级别
	LogSuffix        string // 日志文件后缀
	TimeFormat       string // 时间格式化
	FormatType       string // 格式化类型 json text
	MaxAgeHour       int    // 文件最大保存时间
	RotationTimeHour int    // 日志切割时间间隔
}

type Logger struct {
	Log     *logrus.Logger
	Config  LogConf
	logPath string
}

var (
	levelMap = map[string]logrus.Level{
		"panic": logrus.PanicLevel,
		"fatal": logrus.FatalLevel,
		"warn":  logrus.WarnLevel,
		"info":  logrus.InfoLevel,
		"debug": logrus.DebugLevel,
		"trace": logrus.TraceLevel,
	}
	logger *Logger
)

func Setup(conf LogConf) {
	logPath := conf.Path
	if logPath == "" {
		logPath = "runtime/logs"
	}
	if exist := utils.FileIsExist(logPath); !exist {
		if err := os.MkdirAll(logPath, 0775); err != nil {
			log.Fatalf("Config file create Error: %s, %s", logPath, err.Error())
		}
	}
	level, ok := levelMap[strings.ToLower(conf.Level)]
	if !ok {
		level = logrus.DebugLevel
	}

	logs := logrus.New()
	logs.SetLevel(level)
	logger = &Logger{Config: conf, Log: logs, logPath: logPath}
	logger.setHook()
}

func (l *Logger) getFormatter() logrus.Formatter {
	var formatter logrus.Formatter
	timeFormat := l.Config.TimeFormat
	if timeFormat == "" {
		timeFormat = "Y-m-d H:i:s"
	}
	timeFormat = utils.GetTimeFormatStr(timeFormat)
	if l.Config.FormatType == "json" {
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

func (l *Logger) setHook() {
	formatter := l.getFormatter()
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: l.write("debug"),
		logrus.InfoLevel:  l.write("info"),
		logrus.WarnLevel:  l.write("warning"),
		logrus.ErrorLevel: l.write("error"),
		logrus.FatalLevel: l.write("fatal"),
		logrus.PanicLevel: l.write("panic"),
	}, formatter)
	l.Log.AddHook(lfHook)
}

func (l *Logger) write(level string) *rotatelogs.RotateLogs {
	var (
		fileSuffix           = "%Y-%m-%d-" + level + l.Config.LogSuffix
		fileFullPath         = path.Join(l.logPath, fileSuffix)
		maxAge, motationTime = l.Config.MaxAgeHour, l.Config.RotationTimeHour
		options              = make([]rotatelogs.Option, 0)
	)
	if maxAge > 0 {
		// 文件最大保存时间
		options = append(options, rotatelogs.WithMaxAge(time.Hour*time.Duration(maxAge)))
	}
	if motationTime > 0 {
		// 日志切割时间间隔
		options = append(options, rotatelogs.WithMaxAge(time.Hour*time.Duration(motationTime)))
	}
	// 软连接
	// fileRotateLogs.WithLinkName(logPath)
	writer, err := rotatelogs.New(fileFullPath, options...)
	if err != nil {
		log.Fatalf("Config file write Error: %s", err.Error())
	}
	return writer
}
