package ansicolor

import (
	"errors"
	"strconv"
)

// ErrSGRNotFound indicates that the specified SGR could not be found.
// ErrSGREmpty indicates that the SGR name provided is empty.
var (
	ErrSGRNotFound = errors.New("SGR not found")
	ErrSGREmpty    = errors.New("SGR name is empty")
)

// SGRSetter represents a Select Graphic Rendition (SGR) code for terminal text formatting.
type SGRSetter int

// SGRBold represents a bold text style.
// SGRFaint represents a faint or dim text style.
// SGRItalic represents an italic text style.
// SGRUnderline represents an underline text style.
// SGRBlink represents a slow blink text style.
// SGRFastBlink represents a fast blink text style.
// SGRReverse represents a reverse video or inverted text style.
// SGRConceal represents a concealed or hidden text style.
// SGRStrike represents a strikethrough text style.
// SGRDoubleUnderline represents a double underline text style.
const (
	SGRBold SGRSetter = iota + 1
	SGRFaint
	SGRItalic
	SGRUnderline
	SGRBlink
	SGRFastBlink
	SGRReverse
	SGRConceal
	SGRStrike

	SGRDoubleUnderline SGRSetter = iota + 12
)

// String returns a string representation of the SGRSetter if it is valid, consisting of ANSI escape codes.
func (s SGRSetter) String() string {
	if !s.IsValid() {
		return ""
	}
	return StartFormat + s.Short() + EndFormat
}

// Short returns the SGRSetter value as a string representation of its integer value if it is valid,
// otherwise an empty string.
func (s SGRSetter) Short() string {
	if !s.IsValid() {
		return ""
	}
	return strconv.Itoa(int(s))
}

// Name returns the name associated with the SGRSetter instance if it is valid, otherwise it returns an empty string.
func (s SGRSetter) Name() string {
	if !s.IsValid() {
		return ""
	}
	return SGRNameLookup[s]
}

// IsValid checks whether the SGRSetter value is within the defined valid range of SGR attributes.
func (s SGRSetter) IsValid() bool {
	return s >= 1 && s <= 9 || s == 21
}

// GetReset returns the SGRClearer associated with the SGRSetter to reset its effect.
func (s SGRSetter) GetReset() SGRClearer {
	return SGRResetMap[s]
}

// SGRSetterNameLookup is a map that associates SGRSetter values with their corresponding display names as strings.
type SGRSetterNameLookup map[SGRSetter]string

// SGRNameLookup maps SGRSetter constants to their corresponding human-readable string representations for styling text.
var SGRNameLookup = SGRSetterNameLookup{
	SGRBold:            "bold",
	SGRFaint:           "faint",
	SGRItalic:          "italic",
	SGRUnderline:       "underline",
	SGRBlink:           "blink",
	SGRFastBlink:       "fast blink",
	SGRReverse:         "reverse",
	SGRConceal:         "conceal",
	SGRStrike:          "strike",
	SGRDoubleUnderline: "double underline",
}

// MSGRSetterLookup defines a map where keys are string labels and values are of type SGRSetter for text formatting.
type MSGRSetterLookup map[string]SGRSetter

// SGRSetterLookup maps string keys representing SGR attributes to their corresponding SGRSetter constants.
var SGRSetterLookup = MSGRSetterLookup{
	"bold":             SGRBold,
	"faint":            SGRFaint,
	"italic":           SGRItalic,
	"underline":        SGRUnderline,
	"blink":            SGRBlink,
	"fast blink":       SGRFastBlink,
	"reverse":          SGRReverse,
	"conceal":          SGRConceal,
	"strike":           SGRStrike,
	"double underline": SGRDoubleUnderline,
}

// GetSGRSetterFromString retrieves an SGRSetter associated with the given name or returns an error if not
// found or invalid.
func GetSGRSetterFromString(name string) (SGRSetter, error) {
	if name == "" {
		return -1, ErrSGREmpty
	}
	sgr, ok := SGRSetterLookup[name]
	if !ok {
		return -1, ErrSGRNotFound
	}
	return sgr, nil
}

// SGRResets maps an SGRSetter to its corresponding SGRClearer, defining reset counterparts
// for SGR formatting operations.
type SGRResets map[SGRSetter]SGRClearer

// SGRResetMap maps SGRSetter constants to their respective SGRClearer reset counterparts based on ANSI escape codes.
var SGRResetMap = SGRResets{
	SGRBold:            SGRRemoveIntensity,
	SGRFaint:           SGRRemoveIntensity,
	SGRItalic:          SGRRemoveItalic,
	SGRUnderline:       SGRRemoveUnderline,
	SGRBlink:           SGRRemoveBlink,
	SGRFastBlink:       SGRRemoveBlink,
	SGRReverse:         SGRRemoveReverse,
	SGRConceal:         SGRRemoveConceal,
	SGRStrike:          SGRRemoveStrike,
	SGRDoubleUnderline: SGRRemoveUnderline,
}
