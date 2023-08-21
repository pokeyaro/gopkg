// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog

import (
	"context"
	"log/slog"
	"os"
	"time"
)

type defaulter interface {
	Handler() slog.Handler
}

func defaults(def defaulter) defaulter {
	handler := def.Handler()
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}

var defaultHandler = NewHandler(os.Stdout)

// DefaultLogger initializes and returns the default logger.
func DefaultLogger() Logger {
	handle := NewEntry(defaultHandler)
	defaults(handle)
	return Logger(handle)
}

// DefaultClassical initializes and returns the default classical-style logger.
func DefaultClassical() Classical {
	defaultHandler.mode.SetLog(ModeSimplify)
	handle := NewClassic(defaultHandler)
	defaults(handle)
	return Classical(handle)
}

// Default initializes and returns the default logger.
func Default() Logger {
	return DefaultLogger()
}

// Dev returns a configured *Handler suitable for development environments.
// It sets up a new *Handler instance with various options for development.
func Dev() *Handler {
	handler := NewHandler(os.Stdout)
	handler.timeFmt = time.DateTime
	handler.isColorful = true
	handler.colorScale = ColorTheme("arco")
	handler.mode = NewMode().SetLog(ModeDetail)
	handler.Level = LevelDebug
	return handler
}

// Prod returns a configured *Handler suitable for production environments.
// It sets up a new *Handler instance with appropriate options for production.
func Prod() *Handler {
	handler := NewHandler(os.Stdout)
	handler.Level = LevelInfo
	handler.onFatal = func(ctx context.Context, rec slog.Record) error {
		// TODO: add hooks...
		return nil
	}
	return handler
}

// ConsoleHandler returns the default handler for logging to the console.
func ConsoleHandler() *Handler {
	return defaultHandler
}
