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
