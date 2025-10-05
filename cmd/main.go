package main

import (
	"fmt"

	"github.com/novelgitllc/ansicolor"
)

func main() {
	// Create a new Format instance
	formatted := ansicolor.NewFormat().
		WithForeground(ansicolor.FgRed).
		WithBackground(ansicolor.BgBrightWhite).
		WithOption(ansicolor.SGROptBold)

	// Create a new variation of the Format
	underlined := formatted.WithOption(ansicolor.SGROptUnderline)

	// Create a different variation of the Format
	reversed := formatted.WithOption(ansicolor.SGROptReverse)

	// Apply the Format to the console using the Set() method
	formatted.Set()
	fmt.Println("Hello, World!")

	underlined.Set()
	fmt.Println("Goodbye, World!")

	reversed.Set()
	fmt.Println("reversed colors")

	// Clear all styling - retaining colors
	ansicolor.ClearStyles()
	fmt.Println("reset styles")

	// Reset all styling and colors
	ansicolor.Reset()
	fmt.Println("clear all formatting")

	fmt.Println("You can also " + underlined.Wrap("wrap", true) + " text.")
}
