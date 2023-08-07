// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

import (
	"encoding/json"
	"io"
)

// Logger is an interface that defines the logging methods.
type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)

	Json(l Level, args any)
}

// Entry exposes the configuration settings.
type Entry struct {
	level        Level
	trackAbsPath bool
	timeFormat   DateFmt
	enableColors bool
	recordToFile RecordRule
	lc           *logCore
}

// New creates a new instance of Entry with default settings.
func New() *Entry {
	entry := &Entry{
		level:        LevelDebug,
		trackAbsPath: false,
		timeFormat:   FmtTime,
		enableColors: false,
		recordToFile: &FileRecord{
			Record: false,
		},
		lc: &logCore{
			w: io.Discard,
		},
	}
	return entry
}

// SetupDev configures the logger for development environment.
func SetupDev() Logger {
	entry := New()

	entry.
		SetTrackAbsPath(true).
		SetEnableColors(true).
		SetRecordToFile(&FileRecord{
			Record:  true,
			Pos:     "./logs/",
			Trigger: LevelWarn,
		})

	timeFormat, err := entry.timeFormat.ParseTimeFormat()
	if err != nil {
		panic(err.Error())
	}

	entry.lc.setLogger(LevelDebug, timeFormat, entry.enableColors, entry.recordToFile)

	return entry
}

// SetupProd configures the logger for production environment.
func SetupProd() Logger {
	entry := New()

	entry.
		SetLevel(LevelInfo).
		SetTimeFormat(FmtDatetime).
		SetRecordToFile(&FileRecord{
			Record:  true,
			Pos:     "/tmp/logs/",
			Trigger: LevelError,
		})

	timeFormat, err := entry.timeFormat.ParseTimeFormat()
	if err != nil {
		panic(err.Error())
	}

	entry.lc.setLogger(LevelInfo, timeFormat, entry.enableColors, entry.recordToFile)

	return entry
}

// Debug logs messages at Debug level.
func (entry *Entry) Debug(args ...any) {
	entry.lc.logf(entry, LevelDebug, nil, args...)
}

// Info logs messages at Info level.
func (entry *Entry) Info(args ...any) {
	entry.lc.logf(entry, LevelInfo, nil, args...)
}

// Warn logs messages at Warn level.
func (entry *Entry) Warn(args ...any) {
	entry.lc.logf(entry, LevelWarn, nil, args...)
}

// Error logs messages at Error level.
func (entry *Entry) Error(args ...any) {
	entry.lc.logf(entry, LevelError, nil, args...)
}

// Fatal logs messages at Fatal level.
func (entry *Entry) Fatal(args ...any) {
	entry.lc.logf(entry, LevelFatal, nil, args...)
}

// Debugf logs formatted messages at Debug level.
func (entry *Entry) Debugf(format string, args ...any) {
	entry.lc.logf(entry, LevelDebug, &format, args...)
}

// Infof logs formatted messages at Info level.
func (entry *Entry) Infof(format string, args ...any) {
	entry.lc.logf(entry, LevelInfo, &format, args...)
}

// Warnf logs formatted messages at Warn level.
func (entry *Entry) Warnf(format string, args ...any) {
	entry.lc.logf(entry, LevelWarn, &format, args...)
}

// Errorf logs formatted messages at Error level.
func (entry *Entry) Errorf(format string, args ...any) {
	entry.lc.logf(entry, LevelError, &format, args...)
}

// Fatalf logs formatted messages at Fatal level.
func (entry *Entry) Fatalf(format string, args ...any) {
	entry.lc.logf(entry, LevelFatal, &format, args...)
}

// Json logs a JSON formatted message at the specified level.
func (entry *Entry) Json(l Level, args any) {
	bs, err := json.Marshal(args)
	if err != nil {
		entry.Fatal(err.Error())
	}
	entry.lc.logf(entry, l, nil, string(bs))
}
