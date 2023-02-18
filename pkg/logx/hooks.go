package logx

import (
	"sync"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
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
		fields: make(logrus.Fields, 3),
	}
}

func (g *globalFieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (g *globalFieldsHook) Fire(entry *logrus.Entry) error {
	for k, v := range g.fields {
		entry.Data[k] = v
	}
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
