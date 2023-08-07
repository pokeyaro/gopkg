// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package utils

import (
	"path/filepath"
	"runtime"
)

// GetCallTrace retrieves the file and line number of the specified call stack level.
func GetCallTrace(level int) (string, int) {
	if _, file, lineno, ok := runtime.Caller(level); ok {
		return file, lineno
	} else {
		return "", 0
	}
}

// GetCurrentDirectory retrieves the current directory path.
func GetCurrentDirectory() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return dir
}
