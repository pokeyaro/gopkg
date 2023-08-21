// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync"
	"time"
)

// An Option configures a Handler.
type Option interface {
	apply(*Handler)
}

// HandlerFunc represents a function that configures a Handler.
type HandlerFunc func(*Handler)

// HandlerChain is a collection of HandlerFunc functions.
type HandlerChain []HandlerFunc

// apply applies each HandlerFunc in the chain to the provided Handler.
func (chain HandlerChain) apply(h *Handler) {
	for _, fn := range chain {
		fn(h)
	}
}

// HandlerOptions creates a new Handler with the specified options.
func HandlerOptions(funcs ...HandlerFunc) *Handler {
	h := &Handler{
		builtinSort: []string{FieldTime},
		exitCode:    1,
		absPath:     false,
		timeFmt:     time.DateOnly,
		isColorful:  false,
		colorScale:  nil,
		mode:        NewMode(),
		onFatal:     func(ctx context.Context, rec slog.Record) error { return nil },
		Level:       LevelDebug,
		w:           os.Stdout,
		attrs:       []slog.Attr{},
		groups:      []string{},
		mu:          &sync.Mutex{},
	}

	Option(HandlerChain(funcs)).apply(h)

	return h
}

// WithWriter configures a Handler to use the specified writer.
func WithWriter(w io.Writer) HandlerFunc {
	return func(h *Handler) {
		h.w = w
	}
}

// WithBuiltinSort configures a Handler to use the provided sorting order.
func WithBuiltinSort(sorts []string) HandlerFunc {
	return func(h *Handler) {
		h.builtinSort = sorts
	}
}

// WithLogLevel configures a Handler to use the provided log level.
func WithLogLevel(l Level) HandlerFunc {
	return func(h *Handler) {
		h.Level = l
	}
}

// WithExitCode configures a Handler with the specified exit code.
// This code is used when a fatal log occurs.
func WithExitCode(code int) HandlerFunc {
	return func(h *Handler) {
		h.exitCode = code
	}
}

// WithAbsPath configures a Handler to use either absolute or relative paths.
func WithAbsPath(isAbs bool) HandlerFunc {
	return func(h *Handler) {
		h.absPath = isAbs
	}
}

// WithTimeFormat configures a Handler to use the specified time format.
func WithTimeFormat(timeFmt string) HandlerFunc {
	return func(h *Handler) {
		h.timeFmt = timeFmt
	}
}

// WithColorful configures a Handler to use colorful log output if isColorful is true.
func WithColorful(isColorful bool) HandlerFunc {
	return func(h *Handler) {
		h.isColorful = isColorful
	}
}

// WithColorScale configures a Handler to use the specified color scale.
func WithColorScale(colors *ColorScale) HandlerFunc {
	return func(h *Handler) {
		h.colorScale = colors
	}
}

// WithMode configures a Handler to use the specified mode.
func WithMode(mode *Mode) HandlerFunc {
	return func(h *Handler) {
		h.mode = mode
	}
}

// WithFatalHook configures a Handler with a hook for handling fatal log records.
func WithFatalHook(hook func(ctx context.Context, rec slog.Record) error) HandlerFunc {
	return func(h *Handler) {
		h.onFatal = hook
	}
}

// InitLogger initializes a logger with the current Handler configuration.
func (h *Handler) InitLogger() Logger {
	handle := NewEntry(h)
	defaults(handle)
	return Logger(handle)
}

// InitClassical initializes a classical-style logger with the current Handler configuration.
func (h *Handler) InitClassical() Classical {
	handle := NewClassic(h)
	defaults(handle)
	return Classical(handle)
}

// New creates a new logger with the specified Handler options.
func New(funcs ...HandlerFunc) Logger {
	return HandlerOptions(funcs...).InitLogger()
}
