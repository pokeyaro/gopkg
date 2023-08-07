// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package utils

import (
	"sync"
	"unicode/utf8"
)

// buffer is used to store bytes.
// Use simple []byte instead of bytes.Buffer to avoid large dependency.
type buffer []byte

func (b *buffer) Write(p []byte) {
	*b = append(*b, p...)
}

func (b *buffer) WriteString(s string) {
	*b = append(*b, s...)
}

func (b *buffer) WriteByte(c byte) error {
	*b = append(*b, c)
	return nil
}

func (b *buffer) WriteRune(r rune) {
	if r < utf8.RuneSelf {
		*b = append(*b, byte(r))
		return
	}

	buf := *b
	n := len(buf)
	for n+utf8.UTFMax > cap(buf) {
		buf = append(buf, 0)
	}
	w := utf8.EncodeRune(buf[n:n+utf8.UTFMax], r)
	*b = buf[:n+w]
}

// pp is used to store a printer's state and is reused with sync.Pool to avoid allocations.
type pp struct {
	buf buffer
}

var ppFree = sync.Pool{
	New: func() interface{} {
		return new(pp)
	},
}

func newPrinter() *pp {
	p := ppFree.Get().(*pp)
	return p
}

func (p *pp) free() {
	p.buf = p.buf[:0]
	ppFree.Put(p)
}

func (p *pp) doPrintf(format string, vs []string) string {
	end := len(format)
	argNum := 0
	for i := 0; i < end; {
		lasti := i
		for i < end && format[i] != '%' {
			i++
		}
		if i > lasti {
			p.buf.WriteString(format[lasti:i])
		}
		if i >= end {
			// done processing format string
			break
		}

		// Process one verb
		i++
		if i >= end {
			// done processing format string
			break
		}

		c := format[i]
		switch c {
		case 's':
			if argNum < len(vs) {
				p.printArgString(vs[argNum])
			}
			argNum++
		default:
			p.printByte(c)
		}

		i++
	}

	if argNum < len(vs) {
		for _, arg := range vs[argNum:] {
			p.printArgString(arg)
		}
	}

	s := string(p.buf)
	p.free()
	return s
}

func (p *pp) printByte(c byte) {
	p.buf.WriteByte(c)
}

func (p *pp) printArgString(v string) {
	p.buf.WriteString(v)
}

// Sprintf formats a string according to a format specifier and returns the resulting string.
func Sprintf(format string, v ...string) string {
	p := newPrinter()
	s := p.doPrintf(format, v)
	return s
}
