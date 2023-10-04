# Colorizer

Colorizer is a Go package that allows you to easily add colored text to your command-line applications using ANSI color codes.

## Features

- Simple and intuitive API for adding colors to your text.
- Support for foreground and background colors.
- Easily create custom color templates.
- Lightweight and dependency-free.

## Installation

You can install the Colorizer package using `go get`:

```bash
go get github.com/KaplanOmr/colorizer
```

## Usage

Here's a quick example of how to use Colorizer:

```go
package main

import (
	"fmt"
	"github.com/KaplanOmr/colorizer"
)

func main() {
	// Create a new Colorizer with red foreground color
	redText := colorizer.New("This text is red", colorizer.RED)

	// Print the colored text
	fmt.Println(redText)
}
```