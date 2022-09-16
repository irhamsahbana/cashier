package bootstrap

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger.SetOutput(file)

	return logger
}
