package ansicolor

import "strings"

// StartFormat defines the starting sequence for terminal color formatting.
// EndFormat defines the ending sequence for terminal color formatting.
const (
	StartFormat = "\033["
	EndFormat   = "m"
)

type Format struct {
	Fg   *FgColor // We want to ensure we avoid 0 values here
	Bg   *BgColor // We want to ensure we avoid 0 values here
	SGRs uint8    // 0 values here is ok
}

func (f *Format) Build() string {
	var b strings.Builder
	b.WriteString(StartFormat)
	if f.Fg == nil {
		b.WriteString(FgDefault.Short())
	} else {
		b.WriteString(f.Fg.Short())
	}
	// Bg to be implemented
	//if f.Bg == nil {
	//	b.WriteString(BgDefault.Short())
	//} else {
	//	b.WriteString(f.Bg.Short())
	//}
	b.WriteString(EndFormat)

	return b.String()
}
