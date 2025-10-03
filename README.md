# ansicolor

Human-friendly constants and helpers for working with ANSI color escape codes in Go.

This tiny library provides:
- Named constants for common ANSI colors (foreground and background), including bright variants.
- A lookup map to get a Color by its string name.
- Convenience utilities to set/reset colors in strings and to strip color codes.

## Installation

Go modules:

```
go get github.com/novelgitllc/ansicolor
```

Import:

```
import "github.com/novelgitllc/ansicolor"
```

## Quick start

```go
package main

import (
    "fmt"
    color "github.com/novelgitllc/ansicolor"
)

func main() {
    // Use constants directly
    fmt.Print(color.BrightBlue) // starts bright blue
    fmt.Println("Hello, world!")
    fmt.Print(color.Reset())    // reset to default

    // Wrap a string and optionally reset
    s := color.Set(color.Red, "danger", true)
    fmt.Println(s) // prints red "danger" then resets

    // Remove color codes
    plain := color.Clear(s)
    fmt.Println(plain) // "danger"

    // Lookup by name (case-sensitive)
    if c, err := color.ColorLookup.GetColorFromString("YellowBackground"); err == nil {
        fmt.Println(color.Set(c, "warning", true))
    }
}
```

## API overview

- Type
  - Color: string alias representing an ANSI escape code.

- Constants (selection)
  - Foreground: Black, Red, Green, Yellow, Blue, Magenta, Cyan, White, Default, BrightBlack, BrightRed, BrightGreen, BrightYellow, BrightBlue, BrightMagenta, BrightCyan
  - Background: BlackBackground, RedBackground, GreenBackground, YellowBackground, BlueBackground, MagentaBackground, CyanBackground, WhiteBackground, DefaultBackground, BrightBlackBackground, BrightRedBackground, BrightGreenBackground, BrightYellowBackground, BrightBlueBackground, BrightMagentaBackground, BrightCyanBackground
  - Special: ResetColor

- Lookup
  - ColorLookup: map[string]Color with keys matching the constant names above (e.g., "Red", "YellowBackground", "BrightBlue").
  - (MColorLookup) GetColorFromString(name string) (Color, error)
    - Returns ResetColor with ErrColorEmpty if name is empty.
    - Returns ResetColor with ErrColorNotFound if name is not present.

- Utilities
  - (Color) String() string: returns the raw ANSI sequence.
  - Set(color Color, s string, reset bool) string: prefixes s with color; appends ResetColor if reset is true.
  - Clear(s string) string: strips any known color codes.
  - Reset() string: returns the ANSI reset sequence (same as ResetColor.String()).

## Notes

- Lookup is case-sensitive; use the exact constant name (e.g., "Red", not "red").
- The constants in this library include the bold/bright attribute (e.g., "\u001b[1;31m" for Red). Behavior may vary depending on your terminal settings.
- Clear removes codes that are known to the library (those present in ColorLookup).

## Testing

Run the tests with:

```
go test ./...
```

## License

MIT License. See LICENSE for details.
