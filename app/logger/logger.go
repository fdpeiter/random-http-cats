package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

func NewLogger(source string) *Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	entry := logger.WithField("source", source)
	return &Logger{entry}
}

func (l *Logger) GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		l.WithFields(logrus.Fields{
			"status_code": c.Writer.Status(),
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"duration":    duration,
			"client_ip":   c.ClientIP(),
		}).Info()
	}
}
