package gointelowl

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// LoggerParams represents the fields to configure your logger.
type LoggerParams struct {
	File      io.Writer
	Formatter logrus.Formatter
	Level     logrus.Level
}

// IntelOwlLogger represents a logger to be used by the developer.
// IntelOwlLogger implements the Logrus logger.
//
// Logrus docs: https://github.com/sirupsen/logrus
type IntelOwlLogger struct {
	Logger *logrus.Logger
}

// Init initializes the IntelOwlLogger via LoggerParams
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
