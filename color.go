package inkstamp

import (
	//"fmt"
	//"golang.org/x/term"
	//"io"
	"os"
	"github.com/inkstamp/inkstamp/termcolor"
	//"strconv"
	//"strings"
	//"unicode"
)

type ColorToggle struct {
	//EnableColor bool
	ColorCap termcolor.ColorCap
	EnableColor bool   
}

//====================================
// TOGGLE OPTIONS
//===================================
type ToggleOption func(*toggleConfig)

type  toggleConfig struct {
	stream *os.File
	sniffFlags bool
	forceColor *bool
	//detectCI  bool
}

func defaultToggleConfig() toggleConfig {
	return toggleConfig{
		stream: os.Stdout,
		sniffFlags: true,
		//detectCI: true
	}
}

func Stream(s *os.File) ToggleOption{
	return func(c *toggleConfig) {
		c.stream = s 
	}
}

func FlagToggle(sniff bool) ToggleOption {
	return func(c *toggleConfig) {
		c.sniffFlags = sniff
	}
}

func ForceColor(enabled bool) ToggleOption {
	return func(c *toggleConfig) {
		c.forceColor = &enabled
	}
}


//=============================
// COLOR TOGGLE
//=============================


func defaultConfig() toggleConfig {
	return toggleConfig{
		//Turn off flags sniffing by default and use, os.Stdout
		stream: os.Stdout,
		sniffFlags: false,
		//detectCI: true
	}
}

// Should auto detect tty by default
func NewColorToggle(opts ...ToggleOption) *ColorToggle {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
    
	//Force color precedence over everything
	if cfg.forceColor != nil {
		// when forcecolor is set to false
		if !*cfg.forceColor {
			return &ColorToggle{
				ColorCap: termcolor.ColorNone,
				EnableColor: false,
			}
		}
		//Check terminal capability
	    colorCap := termcolor.Capability(
		    termcolor.Stream(cfg.stream),
			//flags shouldn't override
		    termcolor.FlagToggle(false),
	    )
		//If capability is dumb, give it color 16 as color was forced
		if colorCap == termcolor.ColorNone{
			colorCap = termcolor.Color16
		}
		return &ColorToggle{
			ColorCap: colorCap,
			EnableColor: true,
		}
	}

	//Check terminal capability
	colorCap := termcolor.Capability(
		termcolor.Stream(cfg.stream),
		termcolor.FlagToggle(cfg.sniffFlags),
	)
	
	return &ColorToggle{
		ColorCap: colorCap,
		EnableColor: true,
	}
}


//=============================
// PARSE
//=============================

// Parse template string using auto-detected color support
// User don't need to explicitly define color toggle
func Parse(input string) CompiledTemplate {
	return NewColorToggle().Parse(input)
}

// Parse template string using the toggle's color setting
func (toggle *ColorToggle) Parse(input string) CompiledTemplate {
	if toggle == nil {
		toggle = NewColorToggle()
	}
	parts, currentText := parseLoop(input, toggle.EnableColor, toggle.ColorCap)

	//flush any remainig text
	parts = flushText(parts, currentText)

	return CompiledTemplate{
		Parts:       parts,
		TotalLength: len(input),
	}
}

