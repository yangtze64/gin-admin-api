package logx

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type globalFieldsHook struct {
	fields logrus.Fields
	mut    sync.Mutex
}

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
