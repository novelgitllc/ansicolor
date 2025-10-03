package main

import (
	"fmt"

	"github.com/novelgitllc/ansicolor"
)

func main() {
	err := ansicolor.SetFgColor(ansicolor.FgBlack)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello World")
	ansicolor.ResetFgColor()
	err = ansicolor.Set(ansicolor.Black)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello World")
}
