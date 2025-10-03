package ansicolor

import (
	"testing"
)

func TestColor_String(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		expected string
	}{
		{
			name:     "Bright Blue",
			color:    BrightBlue,
			expected: "\033[1;94m",
		},
		{
			name:     "Red",
			color:    Red,
			expected: "\033[1;31m",
		},
		{
			name:     "Yellow Background",
			color:    YellowBackground,
			expected: "\033[1;43m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.color.String(); got != tt.expected {
				t.Errorf("String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMColorLookup_GetColorFromString(t *testing.T) {

	tests := []struct {
		name     string
		lookup   MColorLookup
		colorStr string
		expected Color
		err      error
	}{
		{
			name:     "Valid color found",
			lookup:   ColorLookup,
			colorStr: "Red",
			expected: Red,
			err:      nil,
		},
		{
			name:     "Color not found",
			lookup:   ColorLookup,
			colorStr: "Purple",
			expected: ResetColor,
			err:      ErrColorNotFound,
		},
		{
			name:     "Empty color string",
			lookup:   ColorLookup,
			colorStr: "",
			expected: ResetColor,
			err:      ErrColorEmpty,
		},
		{
			name:     "Case sensitivity mismatch",
			lookup:   ColorLookup,
			colorStr: "red",
			expected: ResetColor,
			err:      ErrColorNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.lookup.GetColorFromString(tt.colorStr)
			if got != tt.expected {
				t.Errorf("GetColorFromString(%q) = %v, want %v", tt.colorStr, got, tt.expected)
			}
			if (err != nil && tt.err == nil) || (err == nil && tt.err != nil) || (err != nil && err.Error() != tt.err.Error()) {
				t.Errorf("GetColorFromString(%q) error = %v, want %v", tt.colorStr, err, tt.err)
			}
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		s        string
		reset    bool
		expected string
	}{
		{
			name:     "Set color",
			color:    BrightBlue,
			s:        "Hello World",
			reset:    true,
			expected: "\033[1;94mHello World\033[0m",
		},
		{
			name:     "Set color without reset",
			color:    BrightBlue,
			s:        "Hello World",
			reset:    false,
			expected: "\033[1;94mHello World",
		},
		{
			name:     "Set color with empty string",
			color:    BrightBlue,
			s:        "",
			reset:    true,
			expected: "\033[1;94m\033[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Set(tt.color, tt.s, tt.reset)
			if got != tt.expected {
				t.Errorf("Set(%s, %s, %v) = %s, want %s", tt.color, tt.s, tt.reset, got, tt.expected)
			}
		})
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "Clear color",
			s:        "\033[1;94mHello World\033[0m",
			expected: "Hello World",
		},
		{
			name:     "Clear color with empty string",
			s:        "\033[1;94m\033[0m",
			expected: "",
		},
		{
			name:     "Clear color with no color codes",
			s:        "Hello World",
			expected: "Hello World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Clear(tt.s)
			if got != tt.expected {
				t.Errorf("Clear(%s) = %s, want %s", tt.s, got, tt.expected)
			}
		})
	}
}

func TestReset(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "Reset color",
			expected: "\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reset()
			if got != tt.expected {
				t.Errorf("Reset() = %s, want %s", got, tt.expected)
			}
		})
	}
}
