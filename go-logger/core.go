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
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/pokeyaro/gopkg/go-logger/utils"
)

// logCore represents the core parameters.
type logCore struct {
	logger *log.Logger
}

// setLogger sets the logger with the specified log level, date format, and color options.
func (lc *logCore) setLogger(l Level, dt int, isEnableColor bool) *logCore {
	var defaultLogger *log.Logger
	if isEnableColor {
		prefixBgColor := lc.logRefColor(l, l.String()+colorReset+" ", false, false)
		defaultLogger = log.New(color.Output, prefixBgColor, dt)
	} else {
		defaultLogger = log.New(color.Output, l.String()+" ", dt)
	}
	lc.logger = defaultLogger
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

	lc.setLogger(logLevel, timeFormat, entry.enableColors)

	funcPos := lc.getFuncPos(entry.trackAbsPath)
	if format == nil {
		contentText := "%+v"
		if entry.enableColors {
			contentText = lc.logRefColor(logLevel, "%+v", true, false)
		}
		lc.logger.Printf(funcPos+contentText, fmt.Sprint(args...))
	} else {
		contentText := *format
		if entry.enableColors {
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
