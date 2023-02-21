package logx

import (
	"github.com/rifflock/lfshook"
)

type HookOption func() Hook

var globalFields = newGlobalFieldHook()

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
