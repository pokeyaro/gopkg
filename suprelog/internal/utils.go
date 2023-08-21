// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package internal

import (
	"os"
	"path/filepath"
	"runtime"
)

// GetProjectRoot returns the root directory of the current project.
func GetProjectRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Base(currentDir), nil
}

// GetSourceLocation returns the file path and line number of the
// calling function's source code location.
func GetSourceLocation() (string, int) {
	_, file, line, _ := runtime.Caller(5)
	return file, line
}
