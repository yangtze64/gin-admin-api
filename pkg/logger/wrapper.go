package logger

import (
	"github.com/sirupsen/logrus"
	"log"
)

func Debug(args ...interface{}) {
	GetLogger().Debug(args)
}

func Info(args ...interface{}) {
	GetLogger().Info(args)
}

func Warning(args ...interface{}) {
	GetLogger().Warning(args)
}

func Error(args ...interface{}) {
	GetLogger().Error(args)
}

func Fatal(args ...interface{}) {
	GetLogger().Fatal(args)
}

func Panic(args ...interface{}) {
	GetLogger().Panic(args)
}

func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args)
}

func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args)
}

func Warningf(format string, args ...interface{}) {
	GetLogger().Warningf(format, args)
}
func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args)
}

func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args)
}

func Panicf(format string, args ...interface{}) {
	GetLogger().Panicf(format, args)
}

func GetLogger() *logrus.Logger {
	if logger == nil {
		log.Panicln("Logger Uninitialized")
	}
	return logger.Log
}
