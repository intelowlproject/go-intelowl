package gointelowl

import (
	"log"
	"os"
)

type LoggerParams struct {
	Flags int
	File  *os.File
}

type IntelOwlLogger struct {
	debugLogger *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	panicLogger *log.Logger
}

func (intelOwlLogger *IntelOwlLogger) Init(loggerParams *LoggerParams) {

	var fileLogToSet *os.File
	flagsToSet := 0

	if loggerParams.Flags == 0 {
		flagsToSet = log.LstdFlags | log.Lshortfile
	} else {
		flagsToSet = loggerParams.Flags
	}

	if loggerParams.File == nil {
		fileLogToSet = os.Stdout
	} else {
		fileLogToSet = loggerParams.File
	}

	intelOwlLogger.debugLogger = log.New(fileLogToSet, "DEBUG: ", flagsToSet)
	intelOwlLogger.warnLogger = log.New(fileLogToSet, "WARNING: ", flagsToSet)
	intelOwlLogger.errorLogger = log.New(fileLogToSet, "ERROR: ", flagsToSet)
	intelOwlLogger.panicLogger = log.New(fileLogToSet, "PANIC: ", flagsToSet)

}

func (intelOwlLogger *IntelOwlLogger) Debug(msg interface{}) {
	intelOwlLogger.debugLogger.Println(msg)
}

func (intelOwlLogger *IntelOwlLogger) Warn(msg interface{}) {
	intelOwlLogger.warnLogger.Println(msg)
}

func (intelOwlLogger *IntelOwlLogger) Error(msg interface{}) {
	intelOwlLogger.errorLogger.Println(msg)
}

func (intelOwlLogger *IntelOwlLogger) Panic(msg interface{}) {
	intelOwlLogger.panicLogger.Panicln(msg)
}
