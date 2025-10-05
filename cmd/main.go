package main

import (
	"fmt"

	"github.com/novelgitllc/ansicolor"
)

func main() {
	format := ansicolor.NewFormat()
	fmt.Printf("%q\n", format.String())
	formatted := format.WithForeground(ansicolor.FgRed).WithBackground(ansicolor.BgBrightWhite).WithOption(ansicolor.SGROptBold)
	underlined := formatted.WithOption(ansicolor.SGROptUnderline)
	reversed := formatted.WithOption(ansicolor.SGROptReverse)
	formatted.Set()
	fmt.Println("Hello, World!")
	underlined.Set()
	fmt.Println("Goodbye, World!")
	reversed.Set()
	fmt.Println("reversed colors...")
	ansicolor.DefaultFormat()
	fmt.Println("reset colors...")
}
