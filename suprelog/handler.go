// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/pokeyaro/gopkg/suprelog/internal"
	"github.com/pokeyaro/gopkg/suprelog/internal/buffer"

	"github.com/goccy/go-json"
)

// Handler is a log handler that writes log records to an io.Writer,
// formatting them as line-delimited Basic objects.
type Handler struct {
	// Built-in field sorting order for log records
	builtinSort []string

	// Exit code to use for fatal logs
	exitCode int

	// Flag indicating whether to use absolute or relative file paths
	absPath bool

	// Format string for log timestamp display
	timeFmt string

	// Indicates whether to enable colors in log output
	isColorful bool

	// Color scale for log level formatting
	colorScale *ColorScale

	// Mode configuration for log and type format
	mode *Mode

	// Callback function for handling fatal logs
	onFatal func(ctx context.Context, rec slog.Record) error

	// Log level for the handler
	Level Level

	// Attributes to include in each log record
	attrs []slog.Attr

	// Additional groups to include in each log record
	groups []string

	// Mutex for synchronization
	mu *sync.Mutex

	// Writer to output log records
	w io.Writer
}

// Build-in field keys used in log records.
const (
	FieldTime  = "time"
	FieldLevel = "level"
	FieldPos   = "position"
)

// Common error values for invalid fields and keys.
const (
	badField = "!BAD-BUILD-IN"
	badKey   = "!BAD-KEY"
	badMode  = "!BAD-MODE"
)

// NewHandler creates a BasicHandler that writes to w.
func NewHandler(w io.Writer) *Handler {
	return &Handler{
		w: w,
		builtinSort: []string{
			FieldTime,
			FieldLevel,
			FieldPos,
		},
		exitCode:   1,
		absPath:    false,
		timeFmt:    "2006-01-02 15:04:05.000",
		isColorful: false,
		colorScale: NewColorScale(),
		mode:       NewMode().SetLog(ModeDetail),
		onFatal:    func(ctx context.Context, rec slog.Record) error { return nil },
		Level:      LevelDebug,
		attrs:      []slog.Attr{},
		groups:     []string{},
		mu:         &sync.Mutex{},
	}
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	minLevel := LevelInfo.Level()
	if &h.Level != nil {
		minLevel = h.Level.Level()
	}
	return level >= minLevel
}

// WithAttrs returns a new BasicHandler whose attributes consists
// of h's attributes followed by attrs.
func (h *Handler) WithAttrs(as []slog.Attr) slog.Handler {
	if len(as) == 0 || as == nil {
		return h
	}
	h.attrs = append(h.attrs, as...)
	return h
}

// WithGroup adds the specified name to the list of log groups
// associated with the Handler.
func (h *Handler) WithGroup(name string) slog.Handler {
	h.groups = append(h.groups, name)
	return h
}

// Separator for prefix and content.
const ComponentSep = '|'

// Handle formats the given Record as a Basic object on a single line and writes
// it to the provided io.Writer.
//
// Each invocation of Handle results in a single serialized call to io.Writer.Write.
// It formats the log record's timestamp, level, source location, message,
// attributes, and any additional groups in a specified order.
func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	// Create a handle state to manage formatting and output
	state := h.newHandleState(buffer.New(), ComponentSep)

	// Iterate through the user-input key-value pairs in front using slog.Record.Attrs,
	// and store them in an external map[string]string structure.
	fronts := make(map[string]string, r.NumAttrs())
	iter := func(as slog.Attr) bool {
		// Detect and handle mismatched keys
		if as.Key == badKey {
			value := strconv.Quote(as.Value.String())
			message := fmt.Sprintf("Bad key error, please add the appropriate key value for %s.", value)
			panic(message)
		}
		fronts[as.Key] = as.Value.String()
		return true
	}
	r.Attrs(iter)

	// Iterate through the user-configured built-in sort order
	for idx, item := range h.builtinSort {
		switch item {
		case FieldTime:
			// Display log time
			state.appendTime(r.Time.Format(h.timeFmt))
		case FieldLevel:
			// Display log level
			level := h.Level.parse(r.Level)
			state.appendLevel(level)
		case FieldPos:
			// Display log location
			fileName, lineNumber := internal.GetSourceLocation()

			// Check for invalid use of runtime package
			if strings.Contains(fileName, "runtime") {
				panic("Invalid handler, please use the methods of suprelog.")
			}

			// Format the log position based on handler configuration
			if h.absPath {
				state.appendPosition(fileName, lineNumber)
			} else {
				projectRoot, err := internal.GetProjectRoot()
				if err != nil {
					return fmt.Errorf("Failed to get project root: %v\n", err)
				}
				relPath := projectRoot + strings.Split(fileName, projectRoot)[1]
				state.appendPosition(relPath, lineNumber)
			}
		default:
			// Handle unknown fields with a placeholder
			state.buf.WriteString(badField)
		}

		// Add separator if not the last field
		if idx != len(h.builtinSort)-1 {
			state.buf.WriteByte(' ')
		}
	}

	// Display log message
	state.appendString(r.Message)

	// Display user-defined attributes, if any
	if r.NumAttrs() > 0 {
		state.appendSMap(fronts)
	}

	// Append newline character
	state.buf.WriteByte('\n')

	// Acquire a lock to ensure thread safety
	h.mu.Lock()
	defer h.mu.Unlock()

	// Write formatted log record to the specified writer
	_, err := h.w.Write(*state.buf)

	// Handle fatal logs and exit
	if r.Level == LevelFatal.Level() {
		_ = h.onFatal(ctx, r)
		os.Exit(h.exitCode)
	}

	// Return the error encountered during writing,
	// if any, or nil if writing was successful.
	return err
}

// handleState holds state for a single call to BasicHandler.Handle.
type handleState struct {
	h      *Handler
	buf    *buffer.Buffer
	sep    byte
	prefix *buffer.Buffer
}

func (h *Handler) newHandleState(buf *buffer.Buffer, sep byte) handleState {
	s := handleState{
		h:      h,
		buf:    buf,
		sep:    sep,
		prefix: buffer.New(),
	}
	return s
}

func (s *handleState) appendTime(str string) {
	s.buf.WriteByte('[')
	s.buf.WriteString(str)
	s.buf.WriteByte(']')
}

func (s *handleState) appendLevel(str string) {
	if s.h.isColorful {
		bgLevel := s.getColoredLevel(str)
		s.buf.WriteString(bgLevel)
	} else {
		s.buf.WriteByte('[')
		s.buf.WriteString(str)
		s.buf.WriteByte(']')
	}
}

func (s *handleState) appendPosition(str string, line int) {
	s.buf.WriteString(str)
	s.buf.WriteByte(':')
	s.buf.WritePosInt(line)
}

func (s *handleState) addSeparator() {
	s.buf.WriteByte(' ')
	s.buf.WriteByte(s.sep)
	s.buf.WriteByte(' ')
}

func (s *handleState) appendString(str string) {
	if len(s.h.builtinSort) > 0 {
		s.addSeparator()
	}

	switch s.h.mode.log {
	case ModeSimplify:
		s.buf.WriteString(str)
	case ModeDetail:
		s.buf.WriteString(strconv.Quote("msg"))
		s.buf.WriteByte(':')
		s.buf.WriteString(strconv.Quote(str))
	default:
		s.buf.WriteString(badMode)
	}
}

func (s *handleState) appendSMap(ms map[string]string) {
	s.addSeparator()

	switch s.h.mode.typ {
	case ModeText:
		s.appendKVs(ms)
	case ModeJson:
		s.appendJSON(ms)
	default:
		s.buf.WriteString(badMode)
	}
}

func (s *handleState) appendKVs(ms map[string]string) {
	fnText := func() {
		first := true
		for key, value := range ms {
			if !first {
				s.buf.WriteByte(' ')
			} else {
				first = false
			}
			s.buf.WriteString(fmt.Sprintf("%s=%s", key, value))
		}
	}

	switch s.h.mode.log {
	case ModeSimplify:
		fnText()
	case ModeDetail:
		s.buf.WriteString(strconv.Quote(ModeText))
		s.buf.WriteByte(':')
		s.buf.WriteByte('"')
		fnText()
		s.buf.WriteByte('"')
	default:
		s.buf.WriteString(badMode)
	}
}

func (s *handleState) appendJSON(ms map[string]string) {
	fnJson := func() ([]byte, error) {
		data, err := json.Marshal(ms)
		if err != nil {
			return nil, fmt.Errorf("Error marshaling JSON: %v\n", err)
		}
		return data, nil
	}

	switch s.h.mode.log {
	case ModeDetail:
		s.buf.WriteString(strconv.Quote(ModeJson))
		s.buf.WriteByte(':')
		fallthrough
	case ModeSimplify:
		if data, err := fnJson(); err != nil {
			panic(err)
		} else {
			s.buf.WriteString(string(data))
		}
	default:
		s.buf.WriteString(badMode)
	}
}
