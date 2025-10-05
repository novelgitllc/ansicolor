package ansicolor

import (
	"errors"
	"fmt"
)

var (
	ErrColorNotFound = errors.New("color not found")
	ErrColorEmpty    = errors.New("color name is empty")
)

const ResetAll = "\033[0m"

// Reset outputs the ANSI escape code to reset all color attributes.
func Reset() {
	fmt.Print(resetHelper())
}

func resetHelper() string {
	return ResetAll
}
