package ansicolor

import (
	"strings"
)

// SGROption represents a set of options using bit flags for SGR (Select Graphic Rendition) attributes.
type SGROption uint16

// SGROptBold represents the bold text style option.
// SGROptFaint represents the faint or dim text style option.
// SGROptItalic represents the italic text style option.
// SGROptUnderline represents the underline text style option.
// SGROptBlink represents the standard blink text style option.
// SGROptFastBlink represents the fast blink text style option.
// SGROptReverse represents the reverse video text style option.
// SGROptConceal represents the concealed text style option.
// SGROptStrike represents the strike-through text style option.
// SGROptDoubleUnderline represents the double underline text style option.
const (
	SGROptBold SGROption = 1 << iota
	SGROptFaint
	SGROptItalic
	SGROptUnderline
	SGROptBlink
	SGROptFastBlink
	SGROptReverse
	SGROptConceal
	SGROptStrike
	SGROptDoubleUnderline
)

// Set updates the SGROption by enabling the specified options using a bitwise OR operation.
func (s *SGROption) Set(options SGROption) {
	*s |= options
}

// Clear removes the specified options from the SGROption receiver using a bitwise AND NOT operation.
func (s *SGROption) Clear(options SGROption) {
	*s &^= options
}

// Has checks if the SGROption contains all the bits set in the provided options. Returns true if the condition is met.
func (s *SGROption) Has(options SGROption) bool {
	return *s&options == options
}

// HasAny checks if the SGROption contains any of the bits set in the provided options. Returns true if any bit matches.
func (s *SGROption) HasAny(options SGROption) bool {
	return *s&options != 0
}

// Toggle inverts the specified options in the current SGROption, flipping their state between set and unset.
func (s *SGROption) Toggle(options SGROption) {
	*s ^= options
}

func (s *SGROption) String() string {
	if *s == 0 {
		return ""
	}
	var b strings.Builder
	for opt, setter := range SGROptSetterLookup {
		if s.Has(opt) {
			b.WriteString(setter.Short())
			b.WriteString(";")
		}
	}
	str := b.String()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}

func (s *SGROption) ClearString() string {
	if *s == 0 {
		return ""
	}
	var b strings.Builder
	for opt, clearer := range SGROptClearerLookup {
		if s.Has(opt) {
			b.WriteString(clearer.Short())
			b.WriteString(";")
		}
	}
	str := b.String()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}

// MOptSetterLookup maps SGROption values to their corresponding SGRSetter implementations for terminal text formatting.
type MOptSetterLookup map[SGROption]SGRSetter

// SGROptSetterLookup maps SGROptions to their respective SGRSetters for applying text style transformations.
var SGROptSetterLookup = MOptSetterLookup{
	SGROptBold:            SGRBold,
	SGROptFaint:           SGRFaint,
	SGROptItalic:          SGRItalic,
	SGROptUnderline:       SGRUnderline,
	SGROptBlink:           SGRBlink,
	SGROptFastBlink:       SGRFastBlink,
	SGROptReverse:         SGRReverse,
	SGROptConceal:         SGRConceal,
	SGROptStrike:          SGRStrike,
	SGROptDoubleUnderline: SGRDoubleUnderline,
}

// MOptClearerLookup maps SGROption values to their corresponding clearing SGRClearer codes for SGR attribute management.
type MOptClearerLookup map[SGROption]SGRClearer

// SGROptClearerLookup maps SGROptions to their corresponding SGRClearer values for resetting specific text styles.
var SGROptClearerLookup = MOptClearerLookup{
	SGROptBold:            SGRRemoveIntensity,
	SGROptFaint:           SGRRemoveIntensity,
	SGROptItalic:          SGRRemoveItalic,
	SGROptUnderline:       SGRRemoveUnderline,
	SGROptBlink:           SGRRemoveBlink,
	SGROptFastBlink:       SGRRemoveBlink,
	SGROptReverse:         SGRRemoveReverse,
	SGROptConceal:         SGRRemoveConceal,
	SGROptStrike:          SGRRemoveStrike,
	SGROptDoubleUnderline: SGRRemoveUnderline,
}
