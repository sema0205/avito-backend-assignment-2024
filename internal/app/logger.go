package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

func SetLogrus(level string) {
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrusLevel)
	}

	logrus.SetOutput(os.Stdout)
}
