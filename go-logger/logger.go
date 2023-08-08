// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package logger provides a simple, lightweight logging library for Go.
package logger

import (
	"context"
	"encoding/json"
	"io"
)

// Logger is a logger interface that provides logging function with levels.
type Logger interface {
	Trace(args ...any)
	Debug(args ...any)
	Info(args ...any)
	Notice(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)

	Tracef(format string, args ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Noticef(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)

	CtxTracef(ctx context.Context, format string, args ...any)
	CtxDebugf(ctx context.Context, format string, args ...any)
	CtxInfof(ctx context.Context, format string, args ...any)
	CtxNoticef(ctx context.Context, format string, args ...any)
	CtxWarnf(ctx context.Context, format string, args ...any)
	CtxErrorf(ctx context.Context, format string, args ...any)
	CtxFatalf(ctx context.Context, format string, args ...any)
}

// Entry represents the configuration options for a logger entry.
type Entry struct {
	// The log level for the entry
	level Level

	// Indicates whether to track absolute file paths
	trackAbsPath bool

	// The date format for log timestamps
	timeFormat DateFmt

	// Indicates whether to enable colors in log output
	isColorful bool

	// The rule for logging records to a file
	recordToFile RecordRule

	// The underlying log core instance
	lc *logCore
}

// New creates a new instance of Entry with default settings.
func New() *Entry {
	entry := &Entry{
		level:        LevelTrace,
		trackAbsPath: false,
		timeFormat:   FmtTime,
		isColorful:   false,
		recordToFile: &FileRecord{
			ShouldRec: false,
		},
		lc: &logCore{
			w: io.Discard,
		},
	}
	return entry
}

// SetupDev configures the logger for the development environment and returns an interface type, Logger.
func SetupDev() Logger {
	entry := New()

	entry.
		SetLevel(LevelDebug).
		SetTrackAbsPath(true).
		SetEnableColors(true).
		SetRecordToFile(&FileRecord{
			ShouldRec: true,
			FilePath:  "./logs/",
			Trigger:   LevelWarn,
		})

	timeFormat, err := entry.timeFormat.ParseTimeFormat()
	if err != nil {
		panic(err.Error())
	}

	entry.lc.setLogger(LevelDebug, timeFormat, entry.isColorful, entry.recordToFile)

	return entry
}

// SetupProd configures the logger for the production environment and returns an interface type, Logger.
func SetupProd() Logger {
	entry := New()

	entry.
		SetLevel(LevelInfo).
		SetTimeFormat(FmtDatetime).
		SetRecordToFile(&FileRecord{
			ShouldRec: true,
			FilePath:  "/tmp/logs/",
			Trigger:   LevelError,
		})

	timeFormat, err := entry.timeFormat.ParseTimeFormat()
	if err != nil {
		panic(err.Error())
	}

	entry.lc.setLogger(LevelInfo, timeFormat, entry.isColorful, entry.recordToFile)

	return entry
}

// Trace logs messages at Trace level.
func (entry *Entry) Trace(args ...any) {
	entry.lc.logf(entry, LevelTrace, nil, args...)
}

// Debug logs messages at Debug level.
func (entry *Entry) Debug(args ...any) {
	entry.lc.logf(entry, LevelDebug, nil, args...)
}

// Info logs messages at Info level.
func (entry *Entry) Info(args ...any) {
	entry.lc.logf(entry, LevelInfo, nil, args...)
}

// Notice logs messages at Notice level.
func (entry *Entry) Notice(args ...any) {
	entry.lc.logf(entry, LevelNotice, nil, args...)
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

// Tracef logs formatted messages at Trace level.
func (entry *Entry) Tracef(format string, args ...any) {
	entry.lc.logf(entry, LevelTrace, &format, args...)
}

// Debugf logs formatted messages at Debug level.
func (entry *Entry) Debugf(format string, args ...any) {
	entry.lc.logf(entry, LevelDebug, &format, args...)
}

// Infof logs formatted messages at Info level.
func (entry *Entry) Infof(format string, args ...any) {
	entry.lc.logf(entry, LevelInfo, &format, args...)
}

// Noticef logs formatted messages at Notice level.
func (entry *Entry) Noticef(format string, args ...any) {
	entry.lc.logf(entry, LevelNotice, &format, args...)
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

// CtxTracef logs a formatted trace-level message with context.
func (entry *Entry) CtxTracef(ctx context.Context, format string, args ...any) {
	entry.lc.logf(entry, LevelTrace, &format, args...)
}

// CtxDebugf logs a formatted debug-level message with context.
func (entry *Entry) CtxDebugf(ctx context.Context, format string, args ...any) {
	entry.lc.logf(entry, LevelDebug, &format, args...)
}

// CtxInfof logs a formatted info-level message with context.
func (entry *Entry) CtxInfof(ctx context.Context, format string, args ...any) {
	entry.lc.logf(entry, LevelInfo, &format, args...)
}

// CtxNoticef logs a formatted notice-level message with context.
func (entry *Entry) CtxNoticef(ctx context.Context, format string, args ...any) {
	entry.lc.logf(entry, LevelNotice, &format, args...)
}

// CtxWarnf logs a formatted warn-level message with context.
func (entry *Entry) CtxWarnf(ctx context.Context, format string, args ...any) {
	entry.lc.logf(entry, LevelWarn, &format, args...)
}

// CtxErrorf logs a formatted error-level message with context.
func (entry *Entry) CtxErrorf(ctx context.Context, format string, args ...any) {
	entry.lc.logf(entry, LevelError, &format, args...)
}

// CtxFatalf logs a formatted fatal-level message with context.
func (entry *Entry) CtxFatalf(ctx context.Context, format string, args ...any) {
	entry.lc.logf(entry, LevelFatal, &format, args...)
}

// Json logs a JSON representation of the provided arguments at the specified log level.
// It requires an instance of Entry created with the New() function.
func (entry *Entry) Json(l Level, args any) {
	bs, err := json.Marshal(args)
	if err != nil {
		entry.Fatal(err.Error())
	}
	entry.lc.logf(entry, l, nil, string(bs))
}
