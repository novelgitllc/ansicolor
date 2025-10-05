package ansicolor

import (
	"fmt"
	"strings"
)

// StartFormat defines the starting sequence for terminal color formatting.
// EndFormat defines the ending sequence for terminal color formatting.
const (
	StartFormat = "\033["
	EndFormat   = "m"
)

// defaultFormat is the default Format instance with no foreground, background, or options.
// The cached string representation includes explicit codes for resetting colors and each possible option.
var defaultFormat = NewFormat().WithForeground(FgDefault).WithBackground(BgDefault)

// GetDefaultFormat returns the default Format instance
// This value can be set globally with SetDefault()
func GetDefaultFormat() *Format {
	return defaultFormat
}

// DefaultFormat outputs the string representation of the default terminal text format to the standard output.
// Shorthand for defaultFormat.Set()
func DefaultFormat() {
	fmt.Print(defaultFormat.String())
}

// SetDefault sets the global default format to the provided Format instance.
// If the provided Format is nil, this is a no-op.
func SetDefault(format *Format) {
	if format == nil {
		return
	}
	defaultFormat = format
}

type Format struct {
	// avoid zero values - zero represents the ANSI reset code and removes all formatting
	fg   *FgColor
	bg   *BgColor
	opts SGROption // 0 values here is ok, it signifies no additional options
	fStr string    // Cached string representation of the format
}

// NewFormat creates a new instance of Format with default values for foreground, background, and options.
func NewFormat() *Format {
	return &Format{
		fg:   nil,
		bg:   nil,
		opts: 0,
		fStr: "",
	}
}

// WithForeground creates a new Format instance with the specified foreground color while preserving other properties.
func (f *Format) WithForeground(fg FgColor) *Format {
	nf := &Format{
		fg:   &fg,
		bg:   f.bg,
		opts: f.opts,
	}
	nf.gen()
	return nf
}

// WithBackground returns a new Format instance with the specified background color applied,
// keeping other fields unchanged.
func (f *Format) WithBackground(bg BgColor) *Format {
	nf := &Format{
		fg:   f.fg,
		bg:   &bg,
		opts: f.opts,
	}
	nf.gen()
	return nf
}

// WithOption creates a new Format with the specified SGROption applied without modifying the original instance.
func (f *Format) WithOption(opt SGROption) *Format {
	opts := f.opts
	opts.Set(opt)
	nf := &Format{
		fg:   f.fg,
		bg:   f.bg,
		opts: opts,
	}
	nf.gen()
	return nf
}

// HasOption checks if the Format instance contains the specified SGROption.
// Returns true if all the options are set.
func (f *Format) HasOption(opts SGROption) bool {
	return f.opts.Has(opts)
}

// HasAnyOption checks if the Format instance contains any of the specified SGROptions.
// Returns true if any match is found.
func (f *Format) HasAnyOption(opts SGROption) bool {
	return f.opts.HasAny(opts)
}

func (f *Format) gen() {
	// Build the format string
	var b strings.Builder
	// Get started
	b.WriteString(StartFormat)
	// if we have a fg color, add it
	if f.fg != nil {
		b.WriteString(f.fg.Short())
		// if we have a bg color or options, add a semicolon
		b.WriteString(";")
	}
	// if we have a bg color, add it
	if f.bg != nil {
		b.WriteString(f.bg.Short())
		// if we have options, add a semicolon
		b.WriteString(";")
	}
	// if we have options, add them
	if f.opts != 0 {
		b.WriteString(f.opts.String())
	} else if f.opts == 0 {
		b.WriteString(SGRClearStringShort)
	}
	// End the format string
	b.WriteString(EndFormat)
	// Cache the string
	f.fStr = b.String()
}

func (f *Format) String() string {
	return f.fStr
}

func (f *Format) Set() {
	fmt.Print(f.String())
}

func (f *Format) Reset() {
	Reset()
}

func (f *Format) Wrap(s string, reset bool) string {
	var b strings.Builder
	b.WriteString(f.String())
	b.WriteString(s)
	if reset {
		b.WriteString(ResetAll)
	}
	return b.String()
}

func (f *Format) Clear(s string, set bool) string {
	var b strings.Builder
	b.WriteString(StartFormat)
	b.WriteString(FgDefault.Short())
	b.WriteString(BgDefault.Short())
	b.WriteString(f.opts.ClearString())
	b.WriteString(EndFormat)
	b.WriteString(s)
	if set {
		b.WriteString(f.String())
	}
	return b.String()
}
