// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

type RecordRule interface {
	ShouldRecord() bool  // Whether to store the record or not
	GetPosition() string // Location for storing the record
	GetTrigger() Level   // Triggering level
}

type FileRecord struct {
	Record  bool
	Pos     string
	Trigger Level
}

func (r *FileRecord) ShouldRecord() bool {
	return r.Record
}

func (r *FileRecord) GetPosition() string {
	return r.Pos
}

func (r *FileRecord) GetTrigger() Level {
	return r.Trigger
}
