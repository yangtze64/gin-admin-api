package logx

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"sync"
)

type (
	HookOption       func() Hook
	globalFieldsHook struct {
		fields logrus.Fields
		mut    sync.Mutex
	}
)

var globalFields = newGlobalFieldHook()

func newGlobalFieldHook() *globalFieldsHook {
	return &globalFieldsHook{
		fields: make(logrus.Fields, 6),
	}
}

func (g *globalFieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (g *globalFieldsHook) Fire(entry *logrus.Entry) error {
	data := make(logrus.Fields, len(entry.Data)+len(g.fields))
	for k, v := range g.fields {
		data[k] = v
	}
	entry.Data = data
	return nil
}

func WithGlobalFieldHook() Hook {
	return globalFields
}

func WithLfsHook() Hook {
	formatter := log.getFormatter()
	writerMap := make(lfshook.WriterMap, len(levelMap))
	for k, v := range levelMap {
		writerMap[v] = log.fileRotate(k)
	}
	lfHook := lfshook.NewHook(writerMap, formatter)
	return lfHook
}
