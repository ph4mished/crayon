//  Package termcolor detects terminal color capabilities for Go applications.
//
//  # Overview
//
//  termcolor tells you what colors your terminal can display and at what level or depth.
//  It answers like :
//    - Does my terminal support colors at all?
//    - Can it show 256 colors?
//    - Can it show truecolor?
//    - Did the user explicitly enable or disable color?
//
//  termcolor was originally built as a color detection and color flag parsing foundation 
//  for Inkstamp (https://github.com/inkstamp/inkstamp)
//  
//  # Key Features
//
//    - Detects terminal color capability (none, 16, 256, truecolor)
//    - Automatically parses --color and --no-color flag in GNU and POSIX styles
//    - Respects NO_COLOR and FORCE_COLOR environment variables
//    - Handles --color=auto, --color=tty automatically
//    - CI environment detection (Github Actions, Travis, Gitlab, CircleCI)
//
//  # Quick Start
//
//  package main
// 
//  import (
//      "fmt"
//      "github.com/phamio/termcolor"
//  )
//  
//  func main() {
//      cap, ok := termcolor.Capability()
//      if ok {
//          switch cap {
//          case termcolor.ColorTrue:
//              fmt.Println("Terminal supports truecolor")
//          case termcolor.Color256:
//              fmt.Println("Terminal supports 256 colors")
//          case termcolor.Color16:
//              fmt.Println("Terminal supports 16 colors")
//          case termcolor.ColorNone:
//              fmt.Println("Terminal supports no color")
//          }
//      } else {
//            fmt.Println("Could not detect color support")
//      }
//  }
//
//  # Color Capability Levels
//
//  ColorCap represents the color depth supported by the terminal.
//  Values are the actual color counts for easy comparison:
//
//  ColorNone = 0          // No color support (dumb terminal)
//  Color16   = 16         // Basic ANSI 16 colors
//  Color256  = 256        // 256 color palette
//  ColorTrue = 16777216   // Truecolor (16 million colors)
//
//  Capabilities can be compared directly:
//
//  cap, ok := termcolor.Capability()
//  if ok && cap >= termColor.Color256 {
//      // terminal supports at least 256 colors
//  }
//
//  # Detection Chain
//
//  Capability runs through the following checks in priority order,
//  returning as soon as a conclusive result is found:
//
//    1. FORCE_COLOR environment variable
//    2. NO_COLOR environment variable 
//    3. TTY check — is output a terminal at all
//    4. Flag sniffing — --color, --no-color in GNU and POSIX styles
//    5. CI systems — Github Actions, Travis, Gitlab, 
//
//  # Options
//  Capability accepts options to customize detections behaviour:
//
//  // Check stderr instead of stdout
//  cap, ok := termcolor.Capability(termcolor.Stream(os.Stderr))
//
//  // Disable flag sniffing
//  cap, ok := termcolor.Capability(termcolor.SniffFlagToggle(false),)
//
//  // Combined
//  cap, ok := termcolor.Capability(termcolor.Stream(os.Stderr), termcolor.SniffFlagToggle(false),)
//
// Any bug fixes or issues
// see => https://github.com/inkstamp/inkstamp
//
//
// # License
//
// Apache License 2.0 — see LICENSE file for details.  
package termcolor

