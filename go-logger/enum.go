// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package logger provides a simple, lightweight logging library for Go.
package logger

import (
	"fmt"
	"log"
	"strings"
)

// Level represents the logging levels.
type Level uint8

const (
	_ Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelNotice
	LevelWarn
	LevelError
	LevelFatal
)

var (
	levelMapText = map[Level]string{
		LevelTrace:  "trace",
		LevelDebug:  "debug",
		LevelInfo:   "info",
		LevelNotice: "notice",
		LevelWarn:   "warn",
		LevelError:  "error",
		LevelFatal:  "fatal",
	}
)

// String returns the string representation of the log level.
func (l Level) String() string {
	if v, ok := levelMapText[l]; ok {
		return "[" + strings.Title(v) + "]"
	} else {
		return "[?Unknown]"
	}
}

// DateFmt represents the date format options.
type DateFmt float32

const (
	_               = iota // ignore first value by assigning to blank identifier
	FmtTime DateFmt = 1 << iota
	FmtDate
	FmtDatetime
)

// ParseTimeFormat returns the corresponding log time format for DateFmt.
func (f DateFmt) ParseTimeFormat() (int, error) {
	switch f {
	case FmtTime:
		return log.Ltime, nil
	case FmtDate:
		return log.Ldate, nil
	case FmtDatetime:
		return log.LstdFlags, nil
	default:
		return 0, fmt.Errorf("not a valid DateFmt %v", f)
	}
}
