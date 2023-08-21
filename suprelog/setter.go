// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog

import (
	"context"
	"log/slog"

	"github.com/pokeyaro/gopkg/suprelog/internal"
)

// SetLogLevel sets the log level of the handler to the specified level.
func (h *Handler) SetLogLevel(l Level) *Handler {
	h.Level = l
	return h
}

// SetBuiltinSort sets the built-in sort order for log fields.
func (h *Handler) SetBuiltinSort(sorts []string) *Handler {
	h.builtinSort = sorts
	return h
}

// SetTimeFormat sets the time format used for log timestamps.
func (h *Handler) SetTimeFormat(format string) *Handler {
	h.timeFmt = format
	return h
}

// SetColorScale sets the color scale used for log levels.
func (h *Handler) SetColorScale(cs *ColorScale) *Handler {
	h.colorScale = cs
	return h
}

// SetFatalHook sets the fatal hook function to be executed before program exit on fatal logs.
func (h *Handler) SetFatalHook(hook func(ctx context.Context, rec slog.Record) error) *Handler {
	h.onFatal = hook
	return h
}

// ToggleLogPath toggles between using absolute and relative paths in log locations.
func (h *Handler) ToggleLogPath() *Handler {
	internal.Ternary(
		h.absPath,
		func() { h.absPath = false },
		func() { h.absPath = true },
	)
	return h
}

// ToggleLogMode toggles the log mode between detail and simplified.
func (h *Handler) ToggleLogMode() *Handler {
	internal.Ternary(
		h.mode.log == ModeDetail,
		func() { h.mode.log = ModeSimplify },
		func() { h.mode.log = ModeDetail },
	)
	return h
}

// ToggleTypMode toggles the log type mode between TEXT and JSON.
func (h *Handler) ToggleTypMode() *Handler {
	internal.Ternary(
		h.mode.typ == ModeText,
		func() { h.mode.typ = ModeJson },
		func() { h.mode.typ = ModeText },
	)
	return h
}

// ToggleLogColorful toggles whether to use colorful log output.
func (h *Handler) ToggleLogColorful() *Handler {
	internal.Ternary(
		h.isColorful,
		func() { h.isColorful = false },
		func() { h.isColorful = true },
	)
	return h
}
