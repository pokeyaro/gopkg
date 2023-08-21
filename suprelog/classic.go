// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/pokeyaro/gopkg/suprelog/internal/buffer"
)

// Classical interface is inspired by the internal go_logger library.
type Classical interface {
	Trace() Classical
	Debug() Classical
	Info() Classical
	Notice() Classical
	Warn() Classical
	Error() Classical
	Fatal() Classical

	Str(format string, a ...any) Classical
	Int(i int) Classical
	Err(err error) Classical
	Obj(obj any) Classical
	Ctx(ctx context.Context, contextKey string) Classical

	Emit()
}

// Classic represents a logger in classical style,
// inspired by the internal go_logger library.
type Classic struct {
	handler slog.Handler
	buf     *buffer.Buffer
	level   Level
}

// Handler returns the slog handler associated with the Classic logger.
func (c *Classic) Handler() slog.Handler { return c.handler }

// NewClassic creates a new Classic logger instance with the given handler.
func NewClassic(h slog.Handler) *Classic {
	if h == nil {
		panic("nil Handler")
	}
	return &Classic{handler: h}
}

// Level sets the log level for the Classic logger.
func (c *Classic) Level(l Level) *Classic {
	if l.String() == "UNKNOWN" {
		panic("Unknown log level")
	}
	return &Classic{
		buf:   nil,
		level: l,
	}
}

// Syntactic sugar methods for setting log levels.

// Trace sets the log level to TRACE and returns a Classical instance.
func (c *Classic) Trace() Classical {
	return c.Level(LevelTrace)
}

// Debug sets the log level to DEBUG and returns a Classical instance.
func (c *Classic) Debug() Classical {
	return c.Level(LevelDebug)
}

// Info sets the log level to INFO and returns a Classical instance.
func (c *Classic) Info() Classical {
	return c.Level(LevelInfo)
}

// Notice sets the log level to NOTICE and returns a Classical instance.
func (c *Classic) Notice() Classical {
	return c.Level(LevelNotice)
}

// Warn sets the log level to WARN and returns a Classical instance.
func (c *Classic) Warn() Classical {
	return c.Level(LevelWarn)
}

// Error sets the log level to ERROR and returns a Classical instance.
func (c *Classic) Error() Classical {
	return c.Level(LevelError)
}

// Fatal sets the log level to FATAL and returns a Classical instance.
func (c *Classic) Fatal() Classical {
	return c.Level(LevelFatal)
}

// Chain methods for adding log data.

// Str appends a formatted string to the log message.
func (c *Classic) Str(format string, a ...any) Classical {
	c.delimiter().buf.WriteString(fmt.Sprintf(format, a...))
	return c
}

// Int appends an integer value to the log message.
func (c *Classic) Int(i int) Classical {
	c.delimiter().buf.WriteString(strconv.Itoa(i))
	return c
}

// Err appends an error value to the log message.
func (c *Classic) Err(err error) Classical {
	c.delimiter().buf.WriteError(err)
	return c
}

// Obj appends an arbitrary object to the log message.
func (c *Classic) Obj(obj any) Classical {
	c.delimiter().buf.WriteString(fmt.Sprint(obj))
	return c
}

// Ctx appends the value associated with a context key to the log message.
func (c *Classic) Ctx(ctx context.Context, contextKey string) Classical {
	c.delimiter().buf.WriteString(fmt.Sprintf("%v", ctx.Value(contextKey)))
	return c
}

// Emit sends the log message to the handler.
func (c *Classic) Emit() {
	if c.buf == nil {
		return
	}
	slog.Log(context.Background(), c.level.Level(), c.buf.String())
}

func (c *Classic) delimiter() *Classic {
	if c.buf == nil {
		c.buf = new(buffer.Buffer)
	} else {
		c.buf.WriteString(" - ")
	}
	return c
}
