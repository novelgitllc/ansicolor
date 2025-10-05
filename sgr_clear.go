package ansicolor

import (
	"fmt"
	"strconv"
	"strings"
)

// SGRClearer represents a Select Graphic Rendition (SGR) code for resetting terminal text formatting attributes.
type SGRClearer int

// SGRRemoveIntensity clears the intensity attribute from output using SGR codes.
// SGRRemoveItalic clears the italic attribute from output using SGR codes.
// SGRRemoveUnderline clears the underline attribute from output using SGR codes.
// SGRRemoveBlink clears the blink attribute from output using SGR codes.
// SGRRemoveReverse clears the reverse attribute from output using SGR codes.
// SGRRemoveConceal clears the conceal attribute from output using SGR codes.
// SGRRemoveStrike clears the strike-through attribute from output using SGR codes.
const (
	SGRRemoveIntensity SGRClearer = iota + 22
	SGRRemoveItalic
	SGRRemoveUnderline
	SGRRemoveBlink

	SGRRemoveReverse SGRClearer = iota + 22 + 1
	SGRRemoveConceal
	SGRRemoveStrike
)

const SGRClearStringShort = "22;23;24;25;27;28;29"

// String returns the string representation of SGRClearer, including formatting codes if the value is valid.
func (s SGRClearer) String() string {
	if !s.IsValid() {
		return ""
	}
	return StartFormat + s.Short() + EndFormat
}

// Short returns the numeric string representation of SGRClearer if it is valid, otherwise it returns an empty string.
func (s SGRClearer) Short() string {
	if !s.IsValid() {
		return ""
	}
	return strconv.Itoa(int(s))
}

// Name returns the corresponding name for the SGRClearer instance if valid, otherwise returns an empty string.
func (s SGRClearer) Name() string {
	if !s.IsValid() {
		return ""
	}
	return SGRClearerNameLookup[s]
}

// IsValid checks if the SGRClearer value is within the valid range of [22-25] or [27-29].
func (s SGRClearer) IsValid() bool {
	return s >= 22 && s <= 25 || s >= 27 && s <= 29
}

// ClearAllSGRs constructs and prints the ANSI escape sequence to reset all text formatting attributes in the terminal.
func ClearAllSGRs() {
	var b strings.Builder
	b.WriteString(StartFormat)
	b.WriteString(SGRClearStringShort)
	b.WriteString(EndFormat)
	fmt.Print(b.String())
}

// MSGRClearerNameLookup maps SGRClearer values to their respective string descriptions for easy lookup.
type MSGRClearerNameLookup map[SGRClearer]string

// SGRClearerNameLookup maps SGRClearer constants to their corresponding descriptive string representations.
var SGRClearerNameLookup = MSGRClearerNameLookup{
	SGRRemoveIntensity: "remove intensity",
	SGRRemoveItalic:    "remove italic",
	SGRRemoveUnderline: "remove underline",
	SGRRemoveBlink:     "remove blink",
	SGRRemoveReverse:   "remove reverse",
	SGRRemoveConceal:   "remove conceal",
	SGRRemoveStrike:    "remove strike",
}

// GetSGRClearerFromString retrieves an SGRClearer by its name string, returning an error if the
// name is empty or not found.
func GetSGRClearerFromString(name string) (SGRClearer, error) {
	if name == "" {
		return -1, ErrSGREmpty
	}
	sgr, ok := SGRClearerLookup[name]
	if !ok {
		return -1, ErrSGRNotFound
	}
	return sgr, nil
}

// MSGRClearerLookup represents a mapping between string keys and SGRClearer constants.
type MSGRClearerLookup map[string]SGRClearer

// SGRClearerLookup is a map that associates string keys with corresponding SGRClearer constants for SGR
// attribute removal.
var SGRClearerLookup = MSGRClearerLookup{
	"remove intensity": SGRRemoveIntensity,
	"remove italic":    SGRRemoveItalic,
	"remove underline": SGRRemoveUnderline,
	"remove blink":     SGRRemoveBlink,
	"remove reverse":   SGRRemoveReverse,
	"remove conceal":   SGRRemoveConceal,
	"remove strike":    SGRRemoveStrike,
}
