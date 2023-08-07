// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

func (entry *Entry) SetLevel(l Level) *Entry {
	entry.level = l
	return entry
}

func (entry *Entry) SetTrackAbsPath(isAbs bool) *Entry {
	entry.trackAbsPath = isAbs
	return entry
}

func (entry *Entry) SetTimeFormat(dt DateFmt) *Entry {
	entry.timeFormat = dt
	return entry
}

func (entry *Entry) SetEnableColors(enable bool) *Entry {
	entry.enableColors = enable
	return entry
}
