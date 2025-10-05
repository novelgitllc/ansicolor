package ansicolor

import (
	"fmt"
	"strconv"
	"strings"
)

// BgColor represents a terminal background color as an integer.
// It includes values for default, standard, and bright background colors.
type BgColor int

// BgDefault represents the ANSI code for resetting the background color to the terminal's default.
const BgDefault BgColor = 49

// BgBlack represents the black background ANSI color code (40).
// BgRed represents the red background ANSI color code (41).
// BgGreen represents the green background ANSI color code (42).
// BgYellow represents the yellow background ANSI color code (43).
// BgBlue represents the blue background ANSI color code (44).
// BgMagenta represents the magenta background ANSI color code (45).
// BgCyan represents the cyan background ANSI color code (46).
// BgWhite represents the white background ANSI color code (47).
const (
	BgBlack BgColor = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// BgBrightBlack represents the bright black background color with ANSI escape codes.
// BgBrightRed represents the bright red background color with ANSI escape codes.
// BgBrightGreen represents the bright green background color with ANSI escape codes.
// BgBrightYellow represents the bright yellow background color with ANSI escape codes.
// BgBrightBlue represents the bright blue background color with ANSI escape codes.
// BgBrightMagenta represents the bright magenta background color with ANSI escape codes.
// BgBrightCyan represents the bright cyan background color with ANSI escape codes.
// BgBrightWhite represents the bright white background color with ANSI escape codes.
const (
	BgBrightBlack BgColor = iota + 100
	BgBrightRed
	BgBrightGreen
	BgBrightYellow
	BgBrightBlue
	BgBrightMagenta
	BgBrightCyan
	BgBrightWhite
)

// String returns the ANSI escape code representation of the BgColor if valid, or an empty string if invalid.
func (b BgColor) String() string {
	if !b.IsValid() {
		return ""
	}
	return StartFormat + b.Short() + EndFormat
}

// Short returns the integer value of the BgColor as a string if valid, otherwise returns an empty string.
func (b BgColor) Short() string {
	if !b.IsValid() {
		return ""
	}
	return strconv.Itoa(int(b))
}

// Name returns the name of the background color if valid, otherwise it returns an empty string.
func (b BgColor) Name() string {
	if !b.IsValid() {
		return ""
	}
	return BgColorNameLookup[b]
}

// IsValid checks if the BgColor value is within the valid range of background color codes.
func (b BgColor) IsValid() bool {
	return b >= 40 && b <= 47 || b >= 100 && b <= 107 || b == 49
}

// MBgColorNameLookup is a map that associates BgColor values with their corresponding color name strings.
type MBgColorNameLookup map[BgColor]string

// BgColorNameLookup provides a mapping from BgColor constants to their corresponding color name strings.
var BgColorNameLookup = MBgColorNameLookup{
	BgBlack:         "black",
	BgRed:           "red",
	BgGreen:         "green",
	BgYellow:        "yellow",
	BgBlue:          "blue",
	BgMagenta:       "magenta",
	BgCyan:          "cyan",
	BgWhite:         "white",
	BgBrightBlack:   "bright_black",
	BgBrightRed:     "bright_red",
	BgBrightGreen:   "bright_green",
	BgBrightYellow:  "bright_yellow",
	BgBrightBlue:    "bright_blue",
	BgBrightMagenta: "bright_magenta",
	BgBrightCyan:    "bright_cyan",
	BgBrightWhite:   "bright_white",
}

// MBgColorLookup is a mapping of string keys to BgColor values, used for background color lookup and representation.
type MBgColorLookup map[string]BgColor

// BgColorLookup maps string representations of background colors to their respective BgColor constants.
var BgColorLookup = MBgColorLookup{
	"black":          BgBlack,
	"red":            BgRed,
	"green":          BgGreen,
	"yellow":         BgYellow,
	"blue":           BgBlue,
	"magenta":        BgMagenta,
	"cyan":           BgCyan,
	"white":          BgWhite,
	"bright_black":   BgBrightBlack,
	"bright_red":     BgBrightRed,
	"bright_green":   BgBrightGreen,
	"bright_yellow":  BgBrightYellow,
	"bright_blue":    BgBrightBlue,
	"bright_magenta": BgBrightMagenta,
	"bright_cyan":    BgBrightCyan,
	"bright_white":   BgBrightWhite,
}

// GetBgColorFromString converts a string to a BgColor, returning an error if the string is empty
// or the color is not found.
func GetBgColorFromString(s string) (BgColor, error) {
	if s == "" {
		return -1, ErrColorEmpty
	}
	c, ok := BgColorLookup[s]
	if !ok {
		return -1, ErrColorNotFound
	}
	return c, nil
}

// AddBgColor applies a background color to a given string and optionally resets the color formatting at the end.
// It returns the formatted string or an error if the provided color is invalid.
func AddBgColor(color BgColor, s string, reset bool) (string, error) {
	if !color.IsValid() {
		return s, ErrColorNotFound
	}
	var b strings.Builder
	b.WriteString(color.String())
	b.WriteString(s)
	if reset {
		b.WriteString(BgDefault.String())
	}
	return b.String(), nil
}

// ClearBgColor removes all background color formatting from a given string by replacing matching
// patterns in BgColorLookup.
func ClearBgColor(s string) string {
	for _, color := range BgColorLookup {
		s = strings.ReplaceAll(s, color.String(), "")
	}
	return s
}

// SetBgColor changes the terminal background color to the specified `BgColor`.
// Returns an error if the given color is not valid.
func SetBgColor(color BgColor) error {
	if !color.IsValid() {
		return ErrColorNotFound
	}
	fmt.Print(color.String())
	return nil
}

// ResetBgColor resets the background color to the default terminal setting by printing the appropriate escape sequence.
func ResetBgColor() {
	fmt.Print(BgDefault.String())
}
