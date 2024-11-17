package colorizer

import (
	"fmt"
	"strings"
)

// Color represents an ANSI color code
type Color int

// Define colors as custom type constants for better type safety
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	// Bright/high-intensity colors
	BrightBlack   Color = iota + 90 - 8 // 90
	BrightRed                           // 91
	BrightGreen                         // 92
	BrightYellow                        // 93
	BrightBlue                          // 94
	BrightMagenta                       // 95
	BrightCyan                          // 96
	BrightWhite                         // 97
)

// Background represents an ANSI background color code
type Background Color

// Define background colors as custom type constants
const (
	BgBlack Background = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// Bright/high-intensity background colors
	BgBrightBlack   Background = iota + 100 - 8 // 100
	BgBrightRed                                 // 101
	BgBrightGreen                               // 102
	BgBrightYellow                              // 103
	BgBrightBlue                                // 104
	BgBrightMagenta                             // 105
	BgBrightCyan                                // 106
	BgBrightWhite                               // 107
)

// Common text attributes
const (
	Bold      = 1
	Dim       = 2
	Italic    = 3
	Underline = 4
	Blink     = 5
)

const (
	ansiTemplate = "\033[%sm"
	reset        = "0"
)

// Style represents text styling options
type Style struct {
	colors []int
}

// New creates a new Style with the given color options
func New(options ...interface{}) *Style {
	s := &Style{colors: make([]int, 0, len(options))}
	for _, opt := range options {
		switch v := opt.(type) {
		case Color:
			s.colors = append(s.colors, int(v))
		case Background:
			s.colors = append(s.colors, int(v))
		case int:
			s.colors = append(s.colors, v)
		}
	}
	return s
}

// WithColor adds a text color to the style
func (s *Style) WithColor(c Color) *Style {
	s.colors = append(s.colors, int(c))
	return s
}

// WithBackground adds a background color to the style
func (s *Style) WithBackground(bg Background) *Style {
	s.colors = append(s.colors, int(bg))
	return s
}

// WithAttribute adds a text attribute to the style
func (s *Style) WithAttribute(attr int) *Style {
	s.colors = append(s.colors, attr)
	return s
}

// Paint applies the style to the given text
func (s *Style) Paint(text string) string {
	if len(s.colors) == 0 {
		return text
	}

	args := strings.Trim(strings.Replace(fmt.Sprint(s.colors), " ", ";", -1), "[]")
	return fmt.Sprintf(ansiTemplate+"%s"+ansiTemplate, args, text, reset)
}

// Paint is a helper function for quick styling
func Paint(text string, options ...interface{}) string {
	return New(options...).Paint(text)
}
