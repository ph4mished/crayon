# TermColor
**termcolor** tells you what colors your terminal can display and at what level or depth

**It answers questions like**
- Does my terminal supports colors at all?
- Can it show 256 colors?
- Can it show truecolor


## Features
- Detects terminal color capability (dumb, 16, 256, truecolor)
- Automatically parses `--color` and `--no-color` flags in both GNU and POSIX styles.
- Checks if color has being explicitly enabled or disabled via flags or environment variables.
- Respects NO_COLOR and FORCE_COLOR environment variables.
- Falls back to auto-detection when no color flags are provided
- Detects if output is a TTY.
- Automatically handles when user uses `--color=auto` or `--color=tty` or `--color tty`

---

## Installation
```bash
go get -u github.com/phamio/termcolor
```

---

## Usage

### Check color support

```go
package main
import (
    "fmt"
    "github.com/phamio/termcolor"
)

func main(){
     //Check if terminal supports color
    cap , status := termcolor.Capability()
    if status == termcolor.Detected && cap >= termcolor.Color16 {
        fmt.Println("Hurray! Terminal Understands colors")
    } else if status == termcolor.NotTTy {
        fmt.Println("Output is piped(color disabled)")
    } else if status ==termcolor.ExplicitDisabled {
        fmt.Println("Colors explicitly disabled by user")
    } else if status == termcolor.Uncertain {
        fmt.Println("Could not detect - try setting FORCE_COLOR=1")
    }
}
```

---

### Check capability level
```go
package main
import (
    "fmt"
    "github.com/phamio/termcolor"
)

func main(){
    cap, ok := termcolor.Capability()

    ///Only trust capability if detection was successful
    if !ok {
        fmt.Println("Could not reliably detect color support")
        return
    }

    switch cap {
    case termcolor.ColorTrue:
        //16 million colors
        fmt.Println("Terminal supports true color")
    case termcolor.Color256:
        //256 color palette
        fmt.Println("Terminal supports 256-color palette")
    case termcolor.Color16:
       //Basic ansi colors
        fmt.Println("Terminal supports 16 colors")
    case termcolor.ColorNone:
        //No color, plain output
        fmt.Println("Terminal supports no color (dumb terminal)")
    }
}
```

---

## License
Apache License 2.0 — see [LICENSE](LICENSE) file for details.

## About
termcolor was originally built as the color detection and color flag parsing foundation for [Inkstamp](https://github.com/inkstamp/inkstamp)
