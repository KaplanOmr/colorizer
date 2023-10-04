package colorizer

import (
	"fmt"
	"strconv"
	"strings"
)

// Colorizer struct represents a text colorizer with a template and text.
// It allows you to create colored text using ANSI color codes.
type Colorizer struct {
	Template string
	Text     string
}

// ANSI color code constants
const (
	BLACK = iota + 30
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

// ANSI background color code constants
const (
	BG_BLACK = iota + 40
	BG_RED
	BG_GREEN
	BG_YELLOW
	BG_BLUE
	BG_MAGENTA
	BG_CYAN
	BG_WHITE
)

// ANSI code with a placeholder for color formatting
const _ANSI = "\033[{ARGS}m"

// RESET is the value of the ANSI code to reset colors
const RESET = 0

// New creates a Colorizer and returns colored text with the given options.
func New(text string, options ...int) string {
	var c Colorizer
	return c.Make(text, options)
}

// NewTemplate generates a Colorizer with the given options and returns it.
func NewTemplate(options ...int) Colorizer {
	var c Colorizer
	c.generateTemplate(options)
	return c
}

// NewWithTemplate generates colored text using the provided Colorizer.
func NewWithTemplate(text string, colorizer Colorizer) string {
	colorizer.generateText(text)
	return colorizer.Text
}

// Make calls generateTemplate and generateText with the given parameters and returns colored text.
func (c *Colorizer) Make(text string, options []int) string {
	c.generateTemplate(options)
	c.generateText(text)
	return c.Text
}

// generateANSI generates ANSI code with the given arguments.
func (c *Colorizer) generateANSI(args string) string {
	return strings.Replace(_ANSI, "{ARGS}", args, 1)
}

// generateTemplate generates ANSI codes for color and reset based on the provided options.
func (c *Colorizer) generateTemplate(options []int) {
	colorANSI := c.generateANSI(c.combineOptions(options))
	resetANSI := c.generateANSI(strconv.Itoa(RESET))
	c.Template = fmt.Sprintf("%s{TEXT}%s", colorANSI, resetANSI)
}

// generateText generates colored text using the template.
func (c *Colorizer) generateText(text string) {
	c.Text = strings.Replace(c.Template, "{TEXT}", text, 1)
}

// combineOptions combines options with a separator and prepares them for ANSI code formatting.
func (c *Colorizer) combineOptions(options []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(options), " ", ";", -1), "[]")
}
