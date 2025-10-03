package ansicolor

import (
	"fmt"
	"strconv"
	"strings"
)

// FgColor represents a foreground color used for formatting output in terminal-based applications.
// It is an integer-based type that adheres to specific ranges for standard and bright colors.
type FgColor int

const FgDefault FgColor = 39

// FgBlack represents the black foreground color
// FgRed represents the red foreground color
// FgGreen represents the green foreground color.
// FgYellow represents the yellow foreground color.
// FgBlue represents the blue foreground color.
// FgMagenta represents the magenta foreground color.
// FgCyan represents the cyan foreground color.
// FgWhite represents the white foreground color.
const (
	FgBlack FgColor = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// FgBrightBlack represents the bright black foreground color.
// FgBrightRed represents the bright red foreground color.
// FgBrightGreen represents the bright green foreground color.
// FgBrightYellow represents the bright yellow foreground color.
// FgBrightBlue represents the bright blue foreground color.
// FgBrightMagenta represents the bright magenta foreground color.
// FgBrightCyan represents the bright cyan foreground color.
// FgBrightWhite represents the bright white foreground color.
const (
	FgBrightBlack FgColor = iota + 90
	FgBrightRed
	FgBrightGreen
	FgBrightYellow
	FgBrightBlue
	FgBrightMagenta
	FgBrightCyan
	FgBrightWhite
)

// String returns the ANSI escape sequence representation of the FgColor instance for terminal text formatting.
func (c FgColor) String() string {
	if !c.IsValid() {
		return ""
	}
	return StartFormat + c.Short() + EndFormat
}

// Short converts the FgColor value to its string representation as an integer.
func (c FgColor) Short() string {
	if !c.IsValid() {
		return ""
	}
	return strconv.Itoa(int(c))
}

// Name returns the human-readable name of the FgColor based on the FgColorNameLookup map.
func (c FgColor) Name() string {
	if !c.IsValid() {
		return ""
	}
	return FgColorNameLookup[c]
}

// IsValid checks whether the FgColor value is within the valid range of ANSI foreground color codes.
func (c FgColor) IsValid() bool {
	return c >= 30 && c <= 37 || c >= 90 && c <= 97 || c == 39
}

// MFgColorNameLookup is a map that associates FgColor values with their string representations of color names.
type MFgColorNameLookup map[FgColor]string

// FgColorNameLookup is a mapping of foreground color codes to their respective names as strings.
var FgColorNameLookup = MFgColorNameLookup{
	FgBlack:         "black",
	FgRed:           "red",
	FgGreen:         "green",
	FgYellow:        "yellow",
	FgBlue:          "blue",
	FgMagenta:       "magenta",
	FgCyan:          "cyan",
	FgWhite:         "white",
	FgBrightBlack:   "bright black",
	FgBrightRed:     "bright red",
	FgBrightGreen:   "bright green",
	FgBrightYellow:  "bright yellow",
	FgBrightBlue:    "bright blue",
	FgBrightMagenta: "bright magenta",
	FgBrightCyan:    "bright cyan",
	FgBrightWhite:   "bright white",
}

// MFgColorLookup is a map that associates color names (as strings) with their corresponding FgColor values.
type MFgColorLookup map[string]FgColor

// FgColorLookup is a map linking string color names to their corresponding FgColor values.
var FgColorLookup = MFgColorLookup{
	"black":          FgBlack,
	"red":            FgRed,
	"green":          FgGreen,
	"yellow":         FgYellow,
	"blue":           FgBlue,
	"magenta":        FgMagenta,
	"cyan":           FgCyan,
	"white":          FgWhite,
	"bright black":   FgBrightBlack,
	"bright red":     FgBrightRed,
	"bright green":   FgBrightGreen,
	"bright yellow":  FgBrightYellow,
	"bright blue":    FgBrightBlue,
	"bright magenta": FgBrightMagenta,
	"bright cyan":    FgBrightCyan,
	"bright white":   FgBrightWhite,
}

// GetFgColorFromString retrieves the FgColor value corresponding to the provided string key from FgColorLookup.
func GetFgColorFromString(s string) (FgColor, error) {
	if s == "" {
		return -1, ErrColorEmpty
	}
	c, ok := FgColorLookup[s]
	if !ok {
		return -1, ErrColorNotFound
	}
	return c, nil
}

// AddFgColor applies a foreground color to the given string using ANSI escape codes, with an optional reset flag.
// Returns the formatted string with the color applied, or an error if the color is invalid.
func AddFgColor(color FgColor, s string, reset bool) (string, error) {
	if !color.IsValid() {
		return s, ErrColorNotFound
	}
	var b strings.Builder
	b.WriteString(color.String())
	b.WriteString(s)
	if reset {
		b.WriteString(ResetColor.String())
	}
	return b.String(), nil
}

// ClearFgColor removes all ANSI escape sequences associated with FgColor from the
// provided string and returns the result.
func ClearFgColor(s string) string {
	for _, c := range FgColorLookup {
		s = strings.ReplaceAll(s, c.String(), "")
	}
	return s
}

// SetFgColor sets the foreground color for terminal text if the provided FgColor value is valid.
// Returns ErrColorNotFound if the color is invalid.
func SetFgColor(color FgColor) error {
	if !color.IsValid() {
		return ErrColorNotFound
	}
	fmt.Print(color.String())
	return nil
}

// ResetFgColor resets the foreground color in the terminal to the default color using the ANSI escape code.
func ResetFgColor() {
	fmt.Print(FgDefault.String())
}
