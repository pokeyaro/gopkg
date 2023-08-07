// Package logger provides a simple, lightweight logging library for Go.
//
// It offers a convenient way to log messages at different levels (Debug, Info, Warn, Error, Fatal)
// and supports formatted logging. Additionally, it includes a JSON logging feature for structured logging.
//
// Author: Pokeya
// License: MIT

package logger

import (
	"errors"

	"github.com/fatih/color"
)

// colorFunc represents the color functions for foreground and background colors.
type colorFunc struct {
	bg func(...any) string
	fg func(...any) string
}

const (
	colorReset = "\033[0m"

	fgColorDebug    = color.FgHiCyan
	fgColorInfo     = color.FgHiBlue
	fgColorWarn     = color.FgHiYellow
	fgColorError    = color.FgHiMagenta
	fgColorFatal    = color.FgHiRed
	fgColorDefault  = color.FgHiBlack
	fgColorHttpOk   = color.FgGreen
	fgColorHttpFail = color.FgRed

	bgColorDebug    = color.BgHiCyan
	bgColorInfo     = color.BgHiBlue
	bgColorWarn     = color.BgHiYellow
	bgColorError    = color.BgHiMagenta
	bgColorFatal    = color.BgHiRed
	bgColorDefault  = color.BgHiBlack
	bgColorHttpOk   = color.BgGreen
	bgColorHttpFail = color.BgRed
)

var (
	fgDebug    = color.New(fgColorDebug).SprintFunc()
	fgInfo     = color.New(fgColorInfo).SprintFunc()
	fgWarn     = color.New(fgColorWarn).SprintFunc()
	fgError    = color.New(fgColorError).SprintFunc()
	fgFatal    = color.New(fgColorFatal).SprintFunc()
	fgDefault  = color.New(fgColorDefault).SprintFunc()
	fgHttpOk   = color.New(fgColorHttpOk).SprintFunc()
	fgHttpFail = color.New(fgColorHttpFail).SprintFunc()

	bgDebug    = color.New(bgColorDebug).SprintFunc()
	bgInfo     = color.New(bgColorInfo).SprintFunc()
	bgWarn     = color.New(bgColorWarn).SprintFunc()
	bgError    = color.New(bgColorError).SprintFunc()
	bgFatal    = color.New(bgColorFatal).SprintFunc()
	bgDefault  = color.New(bgColorDefault).SprintFunc()
	bgHttpOk   = color.New(bgColorHttpOk).SprintFunc()
	bgHttpFail = color.New(bgColorHttpFail).SprintFunc()
)

var (
	colorMapFunc = map[string]colorFunc{
		"debug":    colorFunc{bg: bgDebug, fg: fgDebug},
		"info":     colorFunc{bg: bgInfo, fg: fgInfo},
		"warn":     colorFunc{bg: bgWarn, fg: fgWarn},
		"error":    colorFunc{bg: bgError, fg: fgError},
		"fatal":    colorFunc{bg: bgFatal, fg: fgFatal},
		"default":  colorFunc{bg: bgDefault, fg: fgDefault},
		"httpOk":   colorFunc{bg: bgHttpOk, fg: fgHttpOk},
		"httpFail": colorFunc{bg: bgHttpFail, fg: fgHttpFail},
	}
)

// logRefColor is used to choose the appropriate color for logging based on log level and options.
// args[0]: isFontColor, args[1]: isHttpLog
func (lc *logCore) logRefColor(logLevel Level, format string, args ...bool) string {
	if len(args) == 0 || len(args) > 2 {
		panic(errors.New("not a valid args length"))
	}

	var cs colorFunc
	// If it's an HTTP log
	if len(args) == 2 && args[1] {
		switch logLevel {
		case LevelDebug, LevelInfo, LevelWarn:
			cs = lc.colorSelector("httpOK")
		case LevelError, LevelFatal:
			cs = lc.colorSelector("httpFail")
		default:
			cs = lc.colorSelector("default")
		}
	} else {
		// Log base color
		switch logLevel {
		case LevelDebug:
			cs = lc.colorSelector("debug")
		case LevelInfo:
			cs = lc.colorSelector("info")
		case LevelWarn:
			cs = lc.colorSelector("warn")
		case LevelError:
			cs = lc.colorSelector("error")
		case LevelFatal:
			cs = lc.colorSelector("fatal")
		default:
			cs = lc.colorSelector("default")
		}
	}

	// Return the formatted string
	if args[0] {
		return cs.fg(format)
	}
	return cs.bg(format)
}

// colorSelector selects the appropriate color function based on the colorType.
func (lc *logCore) colorSelector(colorType string) colorFunc {
	if fn, ok := colorMapFunc[colorType]; ok {
		return fn
	} else {
		return colorMapFunc["default"]
	}
}
