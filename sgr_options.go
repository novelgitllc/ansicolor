package ansicolor

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
	// Needs SGR implementation to tranlate to ANSI escape codes
	panic("implement me")
}
