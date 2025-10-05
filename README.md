# ansicolor

[![Go Reference](https://pkg.go.dev/badge/github.com/novelgitllc/ansicolor.svg)](https://pkg.go.dev/github.com/novelgitllc/ansicolor)
[![Go Report Card](https://goreportcard.com/badge/github.com/novelgitllc/ansicolor)](https://goreportcard.com/report/github.com/novelgitllc/ansicolor)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go library for terminal ANSI color and text formatting with a type-safe, immutable API.

## Installation

```bash
go get github.com/novelgitllc/ansicolor
```

## Usage

The primary interface is the `Format` type, which uses a builder pattern for composing colors and text styles.

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/novelgitllc/ansicolor"
)

func main() {
    // Create a format with red foreground, bright white background, and bold text
    format := ansicolor.NewFormat().
        WithForeground(ansicolor.FgRed).
        WithBackground(ansicolor.BgBrightWhite).
        WithOption(ansicolor.SGROptBold)
    
    // Apply formatting and print
    format.Set()
    fmt.Println("Hello, World!")
    
    // Reset to defaults
    ansicolor.Reset()
}
```

### Immutable Composition

Each `With*()` method returns a new `Format` instance, leaving the original unchanged:

```go
// Create base format
base := ansicolor.NewFormat().
    WithForeground(ansicolor.FgRed).
    WithOption(ansicolor.SGROptBold)

// Create variations
underlined := base.WithOption(ansicolor.SGROptUnderline)
reversed := base.WithOption(ansicolor.SGROptReverse)

// Apply each independently
underlined.Set()
fmt.Println("Goodbye, World!")

reversed.Set()
fmt.Println("reversed colors")
```

### Text Wrapping

Use `Wrap()` to apply formatting to specific text:

```go
format := ansicolor.NewFormat().
    WithForeground(ansicolor.FgGreen).
    WithOption(ansicolor.SGROptBold)

// Wrap text with automatic reset
text := "You can also " + format.Wrap("wrap", true) + " text."
fmt.Println(text)
```

### Selective Clearing

```go
// Clear only text styles, keeping colors
ansicolor.ClearStyles()

// Clear only colors, keeping text styles  
ansicolor.ClearColor()

// Clear all formatting
ansicolor.ClearAll()

// Reset everything to the default format
ansicolor.Reset()
```

## API Reference

### Colors

**Foreground Colors:**
- Standard: `FgBlack`, `FgRed`, `FgGreen`, `FgYellow`, `FgBlue`, `FgMagenta`, `FgCyan`, `FgWhite`
- Bright: `FgBrightBlack`, `FgBrightRed`, `FgBrightGreen`, `FgBrightYellow`, `FgBrightBlue`, `FgBrightMagenta`, `FgBrightCyan`, `FgBrightWhite`
- Default: `FgDefault`

**Background Colors:**
- Standard: `BgBlack`, `BgRed`, `BgGreen`, `BgYellow`, `BgBlue`, `BgMagenta`, `BgCyan`, `BgWhite`
- Bright: `BgBrightBlack`, `BgBrightRed`, `BgBrightGreen`, `BgBrightYellow`, `BgBrightBlue`, `BgBrightMagenta`, `BgBrightCyan`, `BgBrightWhite`
- Default: `BgDefault`

### Text Styles

- `SGROptBold` - Bold text
- `SGROptFaint` - Faint/dim text
- `SGROptItalic` - Italic text
- `SGROptUnderline` - Underlined text
- `SGROptBlink` - Blinking text
- `SGROptFastBlink` - Fast blinking text
- `SGROptReverse` - Reverse video
- `SGROptConceal` - Hidden text
- `SGROptStrike` - Strikethrough text
- `SGROptDoubleUnderline` - Double underlined text

### Format Methods

- `NewFormat()` - Create new Format instance
- `WithForeground(FgColor)` - Set foreground color
- `WithBackground(BgColor)` - Set background color
- `WithOption(SGROption)` - Add text style option
- `Set()` - Apply format to terminal
- `String()` - Get ANSI escape sequence
- `Wrap(string, bool)` - Wrap text with formatting
- `Reset()` - Reset format to defaults

### Global Functions

- `Reset()` - Reset terminal to default format
- `ClearColor()` - Clear foreground and background colors
- `ClearStyles()` - Clear text formatting styles
- `ClearAll()` - Reset all formatting

## License

MIT License - see the [LICENSE](LICENSE) file for details.