package logger

import (
	"fmt"
	"os"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

func New() *Logger {
	log := logrus.New()

	wd, err := os.Getwd()
	if err != nil {
		log.Info(err)
	}

	dir := fmt.Sprintf("%s/artefacts/", wd)

	hook := lfshook.NewHook(
		lfshook.PathMap{
			logrus.InfoLevel:  fmt.Sprintf("%s/info.log", dir),
			logrus.DebugLevel: fmt.Sprintf("%s/debug.log", dir),
			logrus.WarnLevel:  fmt.Sprintf("%s/warn.log", dir),
			logrus.ErrorLevel: fmt.Sprintf("%s/error.log", dir),
			logrus.FatalLevel: fmt.Sprintf("%s/fatal.log", dir),
			logrus.PanicLevel: fmt.Sprintf("%s/panic.log", dir),
		},
		&logrus.JSONFormatter{},
	)

	log.Hooks.Add(hook)

	return &Logger{log}
}

func (l *Logger) Info(msg string) {
	l.log.Info(msg)
}

func (l *Logger) Fatal(err error) {
	l.log.Fatal(err)
}
