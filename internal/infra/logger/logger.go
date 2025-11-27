package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

func New() *Logger {
	log := logrus.New()
	return &Logger{log}
}

func (l *Logger) Info(msg string) {
	l.log.Info(msg)
}

func (l *Logger) Error(err error) {
	l.log.Error(err)
}
