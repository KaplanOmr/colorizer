package colorizer

import (
	"fmt"
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
	_ANSI = "\033[%dm"
	RESET = 0
)

func (c *Colorizer) New(text string, color int) string {
	ansiiColor := c.generateANSII(color)
	ansiiReset := c.generateANSII(RESET)

	return ansiiColor + text + ansiiReset
}

func (c *Colorizer) generateANSII(arg int) string {
	return fmt.Sprintf(_ANSI, arg)
}
