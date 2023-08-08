// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/pokeyaro/gopkg/go-logger/utils"
)

// logCore represents the core parameters.
type logCore struct {
	logger  *log.Logger
	w       io.Writer
	prefix  string
	context []byte
}

const (
	logFile = "records.log"
)

// setLogger sets up the logger with the specified log level, date format, color options, and file logging configuration.
func (lc *logCore) setLogger(l Level, dt int, isColorful bool, recordToFile RecordRule) *logCore {
	isRecordFile := recordToFile.ShouldRecord()
	triggerLevel := recordToFile.GetTrigger()
	filePath := recordToFile.GetPosition()

	// Only records to file if the log level is equal to or higher than the trigger level
	if isRecordFile && l >= triggerLevel {
		file, err := os.OpenFile(filePath+logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err.Error())
		}

		if l >= LevelError {
			lc.w = io.MultiWriter(file, os.Stderr)
		} else {
			lc.w = io.MultiWriter(file, os.Stdout)
		}
	} else {
		// Only terminal output
		if isColorful {
			lc.w = color.Output
		} else {
			lc.w = os.Stdout
		}
	}

	// Define prefix
	if isColorful {
		lc.prefix = lc.logRefColor(l, l.String()+colorReset+" ", false, false)
	} else {
		lc.prefix = l.String() + " "
	}

	// Default Logger
	lc.logger = log.New(lc.w, lc.prefix, dt)

	return lc
}

// logf outputs the log message with the specified log level, format, and arguments.
func (lc *logCore) logf(entry *Entry, logLevel Level, format *string, args ...any) {
	if entry.level > logLevel {
		return
	}

	timeFormat, err := entry.timeFormat.ParseTimeFormat()
	if err != nil {
		panic(err.Error())
	}

	lc.setLogger(logLevel, timeFormat, entry.isColorful, entry.recordToFile)

	funcPos := lc.getFuncPos(entry.trackAbsPath)
	if format == nil {
		contentText := "%+v"
		if entry.isColorful {
			contentText = lc.logRefColor(logLevel, "%+v", true, false)
		}
		lc.logger.Printf(funcPos+contentText, fmt.Sprint(args...))
	} else {
		contentText := *format
		if entry.isColorful {
			contentText = lc.logRefColor(logLevel, *format, true, false)
		}
		lc.logger.Printf(funcPos+contentText, args...)
	}

	if logLevel == LevelFatal {
		os.Exit(1)
	}
}

// getFuncPos retrieves the file execution position.
func (lc *logCore) getFuncPos(isAbsPath bool) string {
	file, lineno := utils.GetCallTrace(4)
	if !isAbsPath {
		path := strings.Split(file, "/")
		if len(path) > 2 {
			file = strings.Join(path[len(path)-2:], "/")
		}
	}
	return utils.Sprintf("%s:%s - ", file, strconv.Itoa(lineno))
}
