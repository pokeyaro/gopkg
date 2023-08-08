// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

// RecordRule defines an interface for configuring logging rules.
type RecordRule interface {
	ShouldRecord() bool  // Indicates whether to store the log record or not.
	GetPosition() string // Returns the location for storing the log record.
	GetTrigger() Level   // Returns the triggering level for logging.
}

// FileRecord represents the configuration for file logging.
type FileRecord struct {
	ShouldRec bool   // Indicates whether file logging is enabled or not
	FilePath  string // The file path where logs will be stored
	Trigger   Level  // The trigger level for file logging
}

// ShouldRecord returns whether file logging is enabled or not.
func (fr *FileRecord) ShouldRecord() bool {
	return fr.ShouldRec
}

// GetPosition returns the file path where logs will be stored.
func (fr *FileRecord) GetPosition() string {
	return fr.FilePath
}

// GetTrigger returns the trigger level for file logging.
func (fr *FileRecord) GetTrigger() Level {
	return fr.Trigger
}
