// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

import (
	"testing"
)

func TestEntry_Optional(t *testing.T) {
	log := &Entry{}

	log.Optional(
		WithLevel(LevelDebug),
		WithTrackAbsPath(true),
		WithTimeFormat(FmtDatetime),
		WithEnableColors(true),
	)

	// TBD...
}
