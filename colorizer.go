package colorizer

import (
	"fmt"
	"strconv"
	"strings"
)

type Colorizer struct {
	Template string
	Text     string
}

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

const (
	_ANSI = "\033[{ARGS}m"
	RESET = 0
)

func New(text string, color int) string {
	var c Colorizer

	return c.Make(text, color)
}

func (c *Colorizer) Make(text string, color int) string {
	c.generateTemplate(color)
	c.generateText(text)

	return c.Text
}

func (c *Colorizer) generateANSI(args string) string {
	return strings.Replace(_ANSI, "{ARGS}", args, 1)
}

func (c *Colorizer) generateTemplate(color int) {
	colorANSI := c.generateANSI(strconv.Itoa(color))
	resetANSI := c.generateANSI(strconv.Itoa(RESET))
	c.Template = fmt.Sprintf("%s{TEXT}%s", colorANSI, resetANSI)
}

func (c *Colorizer) generateText(text string) {
	c.Text = strings.Replace(c.Template, "{TEXT}", text, 1)
}
