// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package logger provides a simple, lightweight logging library for Go.
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
