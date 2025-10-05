package ansicolor

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrColorNotFound = errors.New("color not found")
	ErrColorEmpty    = errors.New("color name is empty")
)

const ClearString = "\033[0m"

// Reset sets the terminal color to the default value.
func Reset() {
	fmt.Print(GetDefaultFormat().String())
}

// ClearColor resets both the foreground and background colors of the terminal to their default values.
func ClearColor() {
	ResetFgColor()
	ResetBgColor()
}

// ClearStyles constructs and prints the ANSI escape sequence to reset all text formatting attributes in the terminal.
func ClearStyles() {
	var b strings.Builder
	b.WriteString(StartFormat)
	b.WriteString(SGRClearStringShort)
	b.WriteString(EndFormat)
	fmt.Print(b.String())
}

// ClearAll resets all text formatting attributes and color settings in the terminal to their default states.
func ClearAll() {
	fmt.Print(clearHelper())
}

// clearHelper returns the ANSI escape code for resetting text formatting and color in the terminal.
func clearHelper() string {
	return ClearString
}
