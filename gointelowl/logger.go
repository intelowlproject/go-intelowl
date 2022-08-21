package gointelowl

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type LoggerParams struct {
	File      io.Writer
	Formatter logrus.Formatter
	Level     logrus.Level
}

type IntelOwlLogger struct {
	Logger *logrus.Logger
}

func (intelOwlLogger *IntelOwlLogger) Init(loggerParams *LoggerParams) {
	logger := logrus.New()

	// Where to log the data!
	if loggerParams.File == nil {
		logger.SetOutput(os.Stdout)
	} else {
		logger.Out = loggerParams.File
	}

	if loggerParams.Formatter != nil {
		logger.SetFormatter(loggerParams.Formatter)
	}

	logger.SetLevel(loggerParams.Level)
	intelOwlLogger.Logger = logger
}
