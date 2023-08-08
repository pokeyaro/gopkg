// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package logger provides a simple, lightweight logging library for Go.
package logger

import (
	"testing"
)

func TestNew(t *testing.T) {
	log := New()

	log.SetLevel(LevelDebug).SetTrackAbsPath(true).SetTimeFormat(FmtDatetime).SetEnableColors(true)

	// TBD...
}

func TestSetupDev(t *testing.T) {
	_ = SetupDev()
}

func TestSetupProd(t *testing.T) {
	_ = SetupProd()
}
