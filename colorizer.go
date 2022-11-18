package colorizer

import (
	"fmt"
	"strconv"
	"strings"
)

//Main struct has Text and Template. Firstly creates a Template and then creates colored Text.
type Colorizer struct {
	Template string
	Text     string
}

const (
	//ANSI code variables for color.
	BLACK = iota + 30
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

const (
	//ANSI code variables for background.
	BG_BLACK = iota + 40
	BG_RED
	BG_GREEN
	BG_YELLOW
	BG_BLUE
	BG_MAGENTA
	BG_CYAN
	BG_WHITE
)

const (
	//ANSI code with a placeholder.
	_ANSI = "\033[{ARGS}m"

	//Value of reset of ANSI code.
	RESET = 0
)

//New, creates colorizer and returns colored text.
func New(text string, options ...int) string {
	var c Colorizer

	return c.Make(text, options)
}

//NewTemplate, generate colorizer with given options and returns generated colorizer.
func NewTemplate(options ...int) Colorizer {
	var c Colorizer

	c.generateTemplate(options)

	return c
}

//NewWithTemplate, generate colored text with the given colorizer.
func NewWithTemplate(text string, colorizer Colorizer) string {
	colorizer.generateText(text)

	return colorizer.Text
}

//Make, calls generateTemplate and generateText func with given parameters and then returns colored text.
func (c *Colorizer) Make(text string, options []int) string {
	c.generateTemplate(options)
	c.generateText(text)

	return c.Text
}

//generateANSI, generates ANSI code with given param.
func (c *Colorizer) generateANSI(args string) string {
	return strings.Replace(_ANSI, "{ARGS}", args, 1)
}

//generateTemplate, generates ANSI codes for color and reset.
func (c *Colorizer) generateTemplate(options []int) {
	colorANSI := c.generateANSI(c.combineOptions(options))
	resetANSI := c.generateANSI(strconv.Itoa(RESET))
	c.Template = fmt.Sprintf("%s{TEXT}%s", colorANSI, resetANSI)
}

//generateText, generates colored text with template.
func (c *Colorizer) generateText(text string) {
	c.Text = strings.Replace(c.Template, "{TEXT}", text, 1)
}

//combineOptions, combine options with separator, and prepares for ANSI code.
func (c *Colorizer) combineOptions(options []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(options), " ", ";", -1), "[]")
}
