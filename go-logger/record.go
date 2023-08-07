// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

type recordRule interface {
	shouldRecord() bool  // Whether to store the record or not
	getPosition() string // Location for storing the record
	getTrigger() Level   // Triggering level
}

type fileRecord struct {
	record  bool
	pos     string
	trigger Level
}

func (r *fileRecord) shouldRecord() bool {
	return r.record
}

func (r *fileRecord) getPosition() string {
	return r.pos
}

func (r *fileRecord) getTrigger() Level {
	return r.trigger
}
