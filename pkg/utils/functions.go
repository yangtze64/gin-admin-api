package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
)

func GetLog(c *gin.Context) *logrus.Entry {
	if logz, ok := c.Get("log"); ok {
		switch logz.(type) {
		case *logrus.Entry:
			return logz.(*logrus.Entry)
		case *logrus.Logger:
			return logz.(*logrus.Logger).WithField("Module", "default")
		default:
			log.Panicln("Logger Not Set")
		}
	}
	return nil
}
