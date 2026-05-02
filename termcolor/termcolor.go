package termcolor
//What of detecting capability once and caching it

import (
	"os"
)

type ColorCap int

const (
	unknown   ColorCap = -1
	ColorNone  ColorCap = 0
	Color16 ColorCap = 16
	Color256 ColorCap = 256
	ColorTrue ColorCap = 16777216
)

type Option func(*config)
type  config struct {
	stream *os.File
	sniffFlags bool
}

func defaultConfig() config {
	return config{
		stream: os.Stdout,
		sniffFlags: true,
	}
}

func Stream(s *os.File) Option{
	return func(c *config) {
		c.stream = s 
	}
}

func FlagToggle(sniff bool) Option {
	return func(c *config) {
		c.sniffFlags = sniff
	}
}


func Capability(opts ...Option) ColorCap {
	cachedArgs := os.Args[1:]
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	colorCap := detectColorCap(cachedArgs, cfg)
	if cfg.sniffFlags {
		getColorFlagCap(cachedArgs, colorCap, cfg)
	}
	return colorCap
}





func getColorFlagCap(cachedArgs []string, colorCap ColorCap, cfg config) ColorCap {
	if checkColorToggle(cachedArgs, "color", "=", colorPosArgs){
		if colorCap != ColorNone {
			return colorCap
		}
		//Forced color on
		return Color16
	} 
    
	// Check --color=auto
	if checkColorToggle(cachedArgs, "color", "=", colorAutoArgs){
		return colorCap
	} 
	
	//Check --no-color flags
    if checkColorToggle(cachedArgs, "no-color", "=", colorNegArgs) || checkColorToggle(cachedArgs, "color", "=", colorNegArgs) || checkColorToggle(cachedArgs,"no-color", "=", colorPosArgs) {
	    return ColorNone
	}
	return ColorNone
}