# Colorizer

[![Go Reference](https://pkg.go.dev/badge/github.com/KaplanOmr/colorizer.svg)](https://pkg.go.dev/github.com/KaplanOmr/colorizer)

Colorizer is a Go package that allows you to easily add colored text to your command-line applications using ANSI color codes.

## Features

- Simple and intuitive API for adding colors to your text
- Support for text colors, background colors, and text attributes
- Chainable style creation
- Reusable styles
- Lightweight and dependency-free

## Installation

You can install the Colorizer package using `go get`:

```bash
go get github.com/KaplanOmr/colorizer
```

## Usage

### Quick Styling

```go
package main

import (
	"fmt"
	"github.com/KaplanOmr/colorizer"
)

func main() {
	// Simple color
	fmt.Println(colorizer.Paint("Error!", colorizer.Red))

	// Multiple styles
	fmt.Println(colorizer.Paint("Warning!", 
		colorizer.Yellow,      // text color
		colorizer.BgRed,       // background color
		colorizer.Bold,        // make it bold
	))
}
```

### Creating Reusable Styles

```go
// Create a style using fluent interface
successStyle := colorizer.New().
	WithColor(colorizer.Green).
	WithBackground(colorizer.BgBlack).
	WithAttribute(colorizer.Bold)

// Use the style multiple times
fmt.Println(successStyle.Paint("Success!"))
fmt.Println(successStyle.Paint("All tests passed!"))

// Create a style directly
errorStyle := colorizer.New(
	colorizer.Red,
	colorizer.BgWhite,
	colorizer.Bold,
	colorizer.Underline,
)

fmt.Println(errorStyle.Paint("Fatal error occurred!"))
```

## Available Options

### Text Colors
- `colorizer.Black`
- `colorizer.Red`
- `colorizer.Green`
- `colorizer.Yellow`
- `colorizer.Blue`
- `colorizer.Magenta`
- `colorizer.Cyan`
- `colorizer.White`

### Bright Text Colors
- `colorizer.BrightBlack`
- `colorizer.BrightRed`
- `colorizer.BrightGreen`
- `colorizer.BrightYellow`
- `colorizer.BrightBlue`
- `colorizer.BrightMagenta`
- `colorizer.BrightCyan`
- `colorizer.BrightWhite`

### Background Colors
- `colorizer.BgBlack`
- `colorizer.BgRed`
- `colorizer.BgGreen`
- `colorizer.BgYellow`
- `colorizer.BgBlue`
- `colorizer.BgMagenta`
- `colorizer.BgCyan`
- `colorizer.BgWhite`

### Bright Background Colors
- `colorizer.BgBrightBlack`
- `colorizer.BgBrightRed`
- `colorizer.BgBrightGreen`
- `colorizer.BgBrightYellow`
- `colorizer.BgBrightBlue`
- `colorizer.BgBrightMagenta`
- `colorizer.BgBrightCyan`
- `colorizer.BgBrightWhite`

### Text Attributes
- `colorizer.Bold`
- `colorizer.Dim`
- `colorizer.Italic`
- `colorizer.Underline`
- `colorizer.Blink`

## License

MIT License - see LICENSE file for details