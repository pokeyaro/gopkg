// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog

import (
	"log/slog"
)

// A Level is the importance or severity of a log event.
// The higher the level, the more important or severe the event.
type Level int

// Log levels constants.
const (
	LevelTrace  Level = -8
	LevelDebug  Level = -4
	LevelInfo   Level = 0
	LevelNotice Level = 2
	LevelWarn   Level = 4
	LevelError  Level = 8
	LevelFatal  Level = 12
)

// String returns a human-readable name for the level.
func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelNotice:
		return "NOTICE"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Level returns the log level as a slog.Level.
// It implements the Leveler interface.
func (l Level) Level() slog.Level { return slog.Level(l) }

// Int returns the integer representation of the log level.
func (l Level) Int() int { return int(l) }

func (l Level) parse(level slog.Level) string {
	switch int(level) {
	case LevelTrace.Int():
		return "TRACE"
	case LevelDebug.Int():
		return "DEBUG"
	case LevelInfo.Int():
		return "INFO"
	case LevelNotice.Int():
		return "NOTICE"
	case LevelWarn.Int():
		return "WARN"
	case LevelError.Int():
		return "ERROR"
	case LevelFatal.Int():
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

//var slogLevel = []slog.Level{
//	slog.LevelDebug,
//	slog.LevelInfo,
//	slog.LevelWarn,
//	slog.LevelError,
//}
//
//func equal(level slog.Level) bool {
//	for _, v := range slogLevel {
//		if v.String() == level.String() {
//			return true
//		}
//	}
//	return false
//}
