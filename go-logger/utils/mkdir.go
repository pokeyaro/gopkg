// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package logger provides a simple, lightweight logging library for Go.
package utils

import (
	"os"
)

// Mkdir creates a directory at the specified file path recursively.
func Mkdir(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = os.MkdirAll(filePath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
