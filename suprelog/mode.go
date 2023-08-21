// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog

// Mode represents the configuration options for log and type modes.
type Mode struct {
	log int    // Log mode: 0 for simplified mode, 1 for detailed mode
	typ string // Type mode: "text" for text mode, "json" for JSON mode
}

// Mode constants for log modes.
const (
	ModeSimplify = 0
	ModeDetail   = 1
)

// Mode constants for type modes.
const (
	ModeText = "text"
	ModeJson = "json"
)

// NewMode creates a new instance of Mode with default configuration.
func NewMode() *Mode {
	return &Mode{
		log: ModeSimplify,
		typ: ModeText,
	}
}

func (m *Mode) SetLog(log int) *Mode {
	m.log = log
	return m
}

func (m *Mode) SetTyp(typ string) *Mode {
	m.typ = typ
	return m
}
