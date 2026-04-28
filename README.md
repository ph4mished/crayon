
# Inkstamp

**Inkstamp is a cross-platform terminal presentation library for Go — colors, alignment, padding and reusable templates in on compact DSL.**

**Parse once, stamp many times**


# Installation

```bash
go get -u github.com/inkstamp/inkstamp
```

# Features

- **Multiple Color Systems**: Named colors, hex codes, RGB, 256-color palette.
- **Full ColorFallback Chain**: Automatic downsampling across all color levels — truecolor -> 256 color palette -> ANSI 16 colors.
- **Left, Right and Center Alignment**: Inline padding (declared directly on placeholders), ANSI-aware.
- **Custom Fill Characters**: Dots, dashes, equals — any character you want.
- **Simple Template System**: Parse once, reuse many times.
- **Comprehensive Styles**: Bold, italic, underline, blink, reverse, hidden, strike-through.
- **Granular Resets**: Individual and full reset codes for precise control.
- **No Escapes Needed**: Texts in [] that aren't colors/styles are left as it is.
- **Cross-Platform**: Full color support on Windows (Windows Terminal, cmd, PowerShell), Linux and macOS.
- **Color Toggling**: Respects NO_COLOR environment variable and detects when output is redirected(TTY detection).

---


# Quick Start

```go
package main

import "github.com/inkstamp/inkstamp"

func main() {
    // Colors
    inkstamp.Parse("[fg=blue bg=yellow]Hello in blue on yellow![reset]").Println()
    inkstamp.Parse("[fg=214]Orange from 256-color palette[reset]").Println()
    inkstamp.Parse("[fg=rgb(255,105,180)]Hot pink text[reset]").Println()
    inkstamp.Parse("[fg=#FF5733]Orange hex color[reset]").Println()

    // Styles
    inkstamp.Parse("[bold] Bold text[reset]").Println()
    inkstamp.Parse("[italic strike] italic strikethrough  text only[reset]").Println()
    
    // Reset specific attributes
    inkstamp.Parse("[bold fg=blue]Blue bold text. [bold=reset]No longer bold, but still blue. [fg=reset]No color, but other styles remain[reset]").Println()
}
```

---
  
# Placeholders
Templates use `[0]`, `[1]`, `[2]` as slots for your data:
```go
temp := inkstamp.Parse("[fg=green][0] scored [1] points[reset]")
temp.Println("Alice", 95) //Alice scored 95 points
temp.Println("Bob", 62)  //Bob scored 62 points
```
**The template stays the same, only the data changes**

---



# Padding and Alignment
Inkstamp measure visible text length (ignoring ANSI codes) so colored text aligns correctly.

### Basic Alignment

```go
row := inkstamp.Parse("[fg=cyan][0:<20][fg=yellow][1:>10][reset]")

row.Println("Alice", "admin")
row.Println("Bob", "user")
row.Println("Charlie", "guest")
```

---

### Center Alignment
```go
title := inkstamp.Parse("[fg=cyan bold][0:^40][reset]")
title.Println("Welcome to Inkstamp")
```


### Custom Fill Characters
Use dots, dashes, equals, or any charaters for padding.

```go
//Dot leaders for reports
report := inkstamp.Parse("[fg=yellow][0:<30:.][fg=white][1:>10][reset]")
report.Println("Total Revenue", "$45,321")
report.Println("Net Profit", "$13,041")

// Separator lines
sep := inkstamp.Parse("[fg=cyan][0:^30:-][reset]")
sep.Println("")
```

### CLI Example

```go
package main

import "github.com/inkstamp/inkstamp"

  var(
  usage = inkstamp.Parse("[fg=white]Usage:[reset]   [fg=cyan][0][reset]")
  example = inkstamp.Parse("[fg=white]Example:[reset] [fg=yellow][0][reset]")
  command  = inkstamp.Parse("[fg=green][0:<5][fg=white][1][reset]")
  )

func ShowHelp() {
    usage.Println("kubectl [command] [TYPE] [NAME] [flags]")
    example.Println("kubectl get pods --namespace default")
    command.Println("get", "Display resources")
    command.Println("describe", "Show details")
}

func main(){
	ShowHelp()
}

```

---


## Why Parse Once

Inkstamp templates are designed to be **parsed once and reused many times**. The template string is compiled into an efficient internal representation — this compilation does the heavy lifting once, so every use after that is fast

Parsing a template once and reusing it is ~12x faster than parsing every time

``` go

// GOOD: parse once at startup
var row  = inkstamp.Parse("[fg=cyan][0:<20][fg=yellow][1:>10][reset]")

//Stamp many times in a loop
for _, user := range users {
    row.Println(user.Name, user.Score) //Fast Path — no re-parsing  
}
```

---

# What Inkstamp Is (And Isn't)
**Inkstamp is for**: CLI help menus, log formatters, reports, tables, banners, progress indicators — any structured or repeated terminal output.

**Inkstamp is not**: A TUI framework, layout engine, or interactive UI library. Use BubbleTea or Lip Gloss for that.

# Philosophy:
**Define once, reuse many times, render correctly everywhere.**

## Documentations
- **[Examples](EXAMPLE.md)** — See code examples
- **[Syntax Reference](SYNTAX.md)** — Complete list of colors, styles, resets, and padding options.
- **[Changelog](CHANGELOG.md)** — Release notes and version history

# NOTES
1. Colors automatically fallback based on terminal capability (truecolor -> 256 color palette -> ANSI 16 colors -> none)
2. Padding measure visible text length , ignoring ANSi escape codes.
3. Templates are immutable
4. Always parse once, reuse many times.

# Terminal Limitations (Not Inkstamp's)
1. Terminal Dependency: Colors only work in terminals that support ANSI escape codes. Legacy Windows cmd has limited style support.
3. Style Support: Some styles like blink and double underline are not universally supported across all terminals — behaviour may vary.

---


# Platform Support
- Linux — full support
- macOS — full support
- Windows Terminal — full support
- Windows cmd — full support, limited styles
- Powershell — full support
- Legacy Windows CMD — limited — some styles may not render

---

# Contributing
Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

# License

Apache License 2.0 — see [LICENSE](LICENSE) file for details.

# Acknowledgments
- [mitchellh/colorstring](https://github.com/mitchellh/colorstring) — Bracket syntax inspiration
- ANSI escape code specifications — the foundation everything is built on.
- The Go community — for testing and feedback
- All contributors who have helped improve this library


