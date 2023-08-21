// Copyright 2023 Pokeya. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package suprelog

import (
	"fmt"
	"strconv"
	"strings"
)

// ColorItem defines the color information for a specific log level.
type ColorItem struct {
	Level string // Log level associated with the color.
	RGB   []int  // RGB color values.
	Hex   string // Hexadecimal color code.
}

// ColorScale defines a set of colors to be used in log output based on log levels.
type ColorScale struct {
	IsRGB  bool        // Indicates whether RGB mode is used.
	Colors []ColorItem // List of color information for different log levels.
}

// NewColorScale creates a new instance of ColorScale with default color settings.
func NewColorScale() *ColorScale {
	// Recommended color gradient schemes customized from
	// ChineseColor website: http://zhongguose.com/
	return &ColorScale{
		// Default mode is to use RGB color mode.
		IsRGB: true,
		Colors: []ColorItem{
			{
				// 星灰 Star Gray
				Level: "Trace",
				RGB:   []int{178, 187, 190},
				Hex:   "#b2bbbe",
			},
			{
				// 柏林蓝 Berlin Blue
				Level: "Debug",
				RGB:   []int{18, 107, 174},
				Hex:   "#126bae",
			},
			{
				// 碧青 Jade Green
				Level: "Info",
				RGB:   []int{92, 179, 204},
				Hex:   "#5cb3cc",
			},
			{
				// 蛙绿 Frog Green
				Level: "Notice",
				RGB:   []int{69, 183, 135},
				Hex:   "#45b787",
			},
			{
				// 佛手黄 Buddha's Hand Yellow
				Level: "Warn",
				RGB:   []int{254, 215, 26},
				Hex:   "#fed71a",
			},
			{
				// 海棠红 Crabapple Red
				Level: "Error",
				RGB:   []int{240, 55, 82},
				Hex:   "#f03752",
			},
			{
				// 牵牛紫 Morning Glory Purple
				Level: "Fatal",
				RGB:   []int{104, 23, 82},
				Hex:   "#681752",
			},
		},
	}
}

const (
	arco = "arco"
	ant  = "ant"
	ele  = "element"
)

// ColorTheme creates a new ColorScale instance based on a specified theme.
func ColorTheme(theme string) *ColorScale {
	colorInit := new(ColorScale)

	switch theme {
	case arco:
		// refer to https://arco.design/
		colorInit = &ColorScale{
			IsRGB: true,
			Colors: []ColorItem{
				{
					// arco-btn-gray
					Level: "TRACE",
					RGB:   []int{201, 205, 212},
				},
				{
					// arco-btn-blue
					Level: "DEBUG",
					RGB:   []int{22, 93, 255},
				},
				{
					// arco-btn-cyan
					Level: "INFO",
					RGB:   []int{20, 201, 201},
				},
				{
					// arco-btn-success
					Level: "NOTICE",
					RGB:   []int{0, 180, 42},
				},
				{
					// arco-btn-warn
					Level: "WARN",
					RGB:   []int{255, 125, 0},
				},
				{
					// arco-btn-danger
					Level: "ERROR",
					RGB:   []int{245, 63, 63},
				},
				{
					// arco-btn-magenta
					Level: "FATAL",
					RGB:   []int{245, 49, 157},
				},
			},
		}
	case ant:
		// refer to https://ant.design/
		colorInit = &ColorScale{
			IsRGB: false,
			Colors: []ColorItem{
				{
					// ant-tag-volcano
					Level: "TRACE",
					Hex:   "#d4380d",
				},
				{
					// ant-tag-processing
					Level: "DEBUG",
					Hex:   "#1677ff",
				},
				{
					// ant-tag-cyan
					Level: "INFO",
					Hex:   "#08979c",
				},
				{
					// ant-tag-lime
					Level: "NOTICE",
					Hex:   "#7cb305",
				},
				{
					// ant-tag-warning
					Level: "WARN",
					Hex:   "#faad14",
				},
				{
					// ant-tag-error
					Level: "ERROR",
					Hex:   "#ff4d4f",
				},
				{
					// ant-tag-purple
					Level: "FATAL",
					Hex:   "#531dab",
				},
			},
		}
	case ele:
		// refer to https://element-plus.org/
		colorInit = &ColorScale{
			IsRGB: false,
			Colors: []ColorItem{
				{
					Level: "TRACE",
					Hex:   "#FFFFFF",
				},
				{
					Level: "DEBUG",
					Hex:   "#409EFF",
				},
				{
					Level: "INFO",
					Hex:   "#79BBFF",
				},
				{
					Level: "NOTICE",
					Hex:   "#67C23A",
				},
				{
					Level: "WARN",
					Hex:   "#E6A23C",
				},
				{
					Level: "ERROR",
					Hex:   "#F56C6C",
				},
				{
					Level: "FATAL",
					Hex:   "#909399",
				},
			},
		}
	default:
		// refer to http://zhongguose.com/
		colorInit = NewColorScale()
	}

	return colorInit
}

// getColoredLevel returns the formatted log level with ANSI color codes.
func (s *handleState) getColoredLevel(level string) string {
	var colorCode string

	scale := s.h.colorScale

	for _, item := range scale.Colors {
		if level == strings.ToUpper(item.Level) {
			if scale.IsRGB {
				colorCode = rgbToAnsi(item.RGB...)
			} else {
				colorCode = hexToAnsi(item.Hex)
			}
			break
		}
	}

	return fmt.Sprintf("%s[%s]\033[0m", colorCode, level)
}

// rgbToAnsi converts RGB color values to ANSI color codes.
func rgbToAnsi(num ...int) string {
	if len(num) != 3 {
		panic("invalid rgb color code")
	}

	r := num[0]
	g := num[1]
	b := num[2]

	closestColor := (r*6/256)*36 + (g*6/256)*6 + (b * 6 / 256)
	return fmt.Sprintf("\033[48;5;%dm", 16+closestColor)
}

// hexToAnsi converts Hexadecimal color codes to ANSI color codes.
func hexToAnsi(hex string) string {
	if len(hex) != 7 || hex[0] != '#' {
		panic("invalid hexadecimal color code")
	}

	r, err := strconv.ParseInt(hex[1:3], 16, 0)
	if err != nil {
		panic(err)
	}
	g, err := strconv.ParseInt(hex[3:5], 16, 0)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseInt(hex[5:7], 16, 0)
	if err != nil {
		panic(err)
	}

	closestColor := (int(r)*6/256)*36 + (int(g)*6/256)*6 + (int(b) * 6 / 256)
	return fmt.Sprintf("\033[48;5;%dm", 16+closestColor)
}
