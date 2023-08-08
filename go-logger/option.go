// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

import (
	"github.com/pokeyaro/gopkg/go-logger/utils"
)

type EntryFunc func(*Entry)
type EntryChain []EntryFunc

func (chains EntryChain) apply(entry *Entry) {
	for _, fn := range chains {
		fn(entry)
	}
}

func (entry *Entry) Optional(funcs ...EntryFunc) Logger {
	EntryChain(funcs).apply(entry)
	return entry
}

func WithLevel(l Level) EntryFunc {
	return func(entry *Entry) {
		entry.level = l
	}
}

func WithTrackAbsPath(isAbsPath bool) EntryFunc {
	return func(entry *Entry) {
		entry.trackAbsPath = isAbsPath
	}
}

func WithTimeFormat(dt DateFmt) EntryFunc {
	return func(entry *Entry) {
		entry.timeFormat = dt
	}
}

func WithEnableColors(isEnabled bool) EntryFunc {
	return func(entry *Entry) {
		entry.isColorful = isEnabled
	}
}

func WithRecordToFile(record RecordRule) EntryFunc {
	return func(entry *Entry) {
		filePath := record.GetPosition()
		if err := utils.Mkdir(filePath); err != nil {
			panic(err.Error())
		}
		entry.recordToFile = record
	}
}
