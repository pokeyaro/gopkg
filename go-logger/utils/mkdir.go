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
