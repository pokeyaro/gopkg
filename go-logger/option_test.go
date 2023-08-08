// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package logger provides a simple, lightweight logging library for Go.
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
