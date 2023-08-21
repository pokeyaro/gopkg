// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog

import (
	"context"
	"fmt"
	"log/slog"
)

// Logger represents a generic logging interface with
// various log levels and context support.
type Logger interface {
	Trace(msg string, args ...any)
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Notice(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)

	Tracef(format string, args ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Noticef(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)

	TraceCtx(ctx context.Context, msg string, args ...any)
	DebugCtx(ctx context.Context, msg string, args ...any)
	InfoCtx(ctx context.Context, msg string, args ...any)
	NoticeCtx(ctx context.Context, msg string, args ...any)
	WarnCtx(ctx context.Context, msg string, args ...any)
	ErrorCtx(ctx context.Context, msg string, args ...any)
	FatalCtx(ctx context.Context, msg string, args ...any)
}

// Entry represents a logger entry for structured logging.
type Entry struct {
	handler slog.Handler // for structured logging
}

// Handler returns slog's Handler.
func (e *Entry) Handler() slog.Handler { return e.handler }

// NewEntry creates a new Entry logger instance with the given handler.
func NewEntry(h slog.Handler) *Entry {
	if h == nil {
		panic("nil Handler")
	}
	return &Entry{h}
}

// Trace logs a trace message using the default logger.
func (e *Entry) Trace(msg string, args ...any) {
	slog.Log(context.Background(), LevelTrace.Level(), msg, args...)
}

// Tracef logs a formatted trace message using the default logger.
func (e *Entry) Tracef(format string, args ...any) {
	slog.Log(context.Background(), LevelTrace.Level(), fmt.Sprintf(format, args...))
}

// TraceCtx logs a trace message using the provided context and default logger.
func (e *Entry) TraceCtx(ctx context.Context, msg string, args ...any) {
	slog.Log(ctx, LevelTrace.Level(), msg, args...)
}

// Debug logs a debug message using the default logger.
func (e *Entry) Debug(msg string, args ...any) {
	slog.Log(context.Background(), LevelDebug.Level(), msg, args...)
}

// Debugf logs a formatted debug message using the default logger.
func (e *Entry) Debugf(format string, args ...any) {
	slog.Log(context.Background(), LevelDebug.Level(), fmt.Sprintf(format, args...))
}

// DebugCtx logs a debug message using the provided context and default logger.
func (e *Entry) DebugCtx(ctx context.Context, msg string, args ...any) {
	slog.Log(ctx, LevelDebug.Level(), msg, args...)
}

// Info logs an informational message using the default logger.
func (e *Entry) Info(msg string, args ...any) {
	slog.Log(context.Background(), LevelInfo.Level(), msg, args...)
}

// Infof logs a formatted informational message using the default logger.
func (e *Entry) Infof(format string, args ...any) {
	slog.Log(context.Background(), LevelInfo.Level(), fmt.Sprintf(format, args...))
}

// InfoCtx logs an informational message using the provided context and default logger.
func (e *Entry) InfoCtx(ctx context.Context, msg string, args ...any) {
	slog.Log(ctx, LevelInfo.Level(), msg, args...)
}

// Notice logs a notice message using the default logger.
func (e *Entry) Notice(msg string, args ...any) {
	slog.Log(context.Background(), LevelNotice.Level(), msg, args...)
}

// Noticef logs a formatted notice message using the default logger.
func (e *Entry) Noticef(format string, args ...any) {
	slog.Log(context.Background(), LevelNotice.Level(), fmt.Sprintf(format, args...))
}

// NoticeCtx logs a notice message using the provided context and default logger.
func (e *Entry) NoticeCtx(ctx context.Context, msg string, args ...any) {
	slog.Log(ctx, LevelNotice.Level(), msg, args...)
}

// Warn logs a warning message using the default logger.
func (e *Entry) Warn(msg string, args ...any) {
	slog.Log(context.Background(), LevelWarn.Level(), msg, args...)
}

// Warnf logs a formatted warning message using the default logger.
func (e *Entry) Warnf(format string, args ...any) {
	slog.Log(context.Background(), LevelWarn.Level(), fmt.Sprintf(format, args...))
}

// WarnCtx logs a warning message using the provided context and default logger.
func (e *Entry) WarnCtx(ctx context.Context, msg string, args ...any) {
	slog.Log(ctx, LevelWarn.Level(), msg, args...)
}

// Error logs an error message using the default logger.
func (e *Entry) Error(msg string, args ...any) {
	slog.Log(context.Background(), LevelError.Level(), msg, args...)
}

// Errorf logs a formatted error message using the default logger.
func (e *Entry) Errorf(format string, args ...any) {
	slog.Log(context.Background(), LevelError.Level(), fmt.Sprintf(format, args...))
}

// ErrorCtx logs a error message using the provided context and default logger.
func (e *Entry) ErrorCtx(ctx context.Context, msg string, args ...any) {
	slog.Log(ctx, LevelError.Level(), msg, args...)
}

// Fatal logs a fatal error message using the default logger.
func (e *Entry) Fatal(msg string, args ...any) {
	slog.Log(context.Background(), LevelFatal.Level(), msg, args...)
}

// Fatalf logs a formatted fatal message using the default logger.
func (e *Entry) Fatalf(format string, args ...any) {
	slog.Log(context.Background(), LevelFatal.Level(), fmt.Sprintf(format, args...))
}

// FatalCtx logs a fatal message using the provided context and default logger.
func (e *Entry) FatalCtx(ctx context.Context, msg string, args ...any) {
	slog.Log(ctx, LevelFatal.Level(), msg, args...)
}
