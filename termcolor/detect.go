package termcolor

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"golang.org/x/term"
)

	

func detectColorCap(cachedArgs []string, cfg config) ColorCap {

	//Check FORCE_COLOR level
	if cap, ok := getForceColorLevel(); ok {
		return cap
	}

	//Check NO_COLOR env variable
	if _, exists := os.LookupEnv("NO_COLOR"); exists{
		return ColorNone
	}


	//Check if TTY
	if !term.IsTerminal(int(cfg.stream.Fd())) {
		return ColorNone
	}

	//Stream is a tty, now determine level  
	//CI Systems
	    if cap, ok := getCIColorLevel(); ok{
		    return cap
	    }

	if cap, ok := getKnownTerminal(); ok{
		return cap
	}
    
	//Check TERM_PROGRAM
	if cap, ok := checkTermProgram(); ok{
		return cap
	}

	//Check TERM
	if cap, ok := checkTermEnv(); ok {
		return cap
	}

	//Query tput
	if cap, ok := queryTput(); ok {
		return cap
	}
	
	return ColorNone
}





//================================
// CHECKING TERMINAL ENVIRONMENT
//================================

func checkTermEnv() (ColorCap, bool) {

	term := os.Getenv("TERM")
	colorTerm := os.Getenv("COLORTERM")
	//Check for explicit color number in $TERM or $COLORTERM
	if supportsTrueColor(colorTerm) {
		return ColorTrue, true
	}
	if supports256Color(term){
		return Color256, true
	}
	if supports16Color(term){
		return Color16, true
	}

	if supportsNone(term){
		return ColorNone, true
	}

	//Detection failed
	return unknown, false
}


//====================
// FORCE_COLOR
//====================

func getForceColorLevel() (ColorCap, bool) {
	value, exists := os.LookupEnv("FORCE_COLOR")
	if !exists{
		return unknown, false
	}
	switch value{
	case "0":
		return ColorNone, true
	case "1":
		return Color16, true
	case "2":
		return Color256, true
	case "3":
		return ColorTrue, true
	default:
		return unknown, false
	}
}


//=================
// CI SYSTEMS
//=================

func getCIColorLevel() (ColorCap, bool) {
	if _, exists := os.LookupEnv("CI"); !exists{
		return unknown, false
	}
	switch {
	case envExists("GITHUB_ACTIONS"), envExists("TRAVIS"), envExists("GITLAB_CI"), envExists("CIRCLECI"):
		return Color16, true
	default:
		return unknown, false
	}
}


//=====================
// KNOWN TERMINALS
//=====================

func getKnownTerminal() (ColorCap, bool) {
	switch {
	case envExists("WT_SESSION"), envExists("KONSOLE_VERSION"), envExists("ITERM_SESSION_ID"), envExists("VTE_VERSION"):
		return ColorTrue, true
	default:
		return unknown, false
	}
}

//====================
// FORCE_COLOR
//====================

func checkTermProgram() (ColorCap, bool) {
	value, exists := os.LookupEnv("TERM_PROGRAM")
	if !exists{
		return unknown, false
	}
	switch value{
	case "iTerm.app", "WezTerm", "Hyper", "rio", "vscode", "tabby":
		return ColorTrue, true
	case "Apple_Terminal":
		return Color256, true
	default:
		return unknown, false
	}
}




//=====================
// QUERYING TPUT
//=====================

func queryTput() (ColorCap, bool) {
	//Check if tput exists
	_, err := exec.LookPath("tput")
	if err != nil {
		//Detection failed
		return unknown, false
	}
    
	//Run tput colors
	cmd := exec.Command("tput", "colors")
	out, err := cmd.Output()
	if err != nil {
		return unknown, false
	}
	
	colors, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil || colors <= 0 {
		return unknown, false
	}

	return ColorCap(colors), true
}





//=======================================
// MINI TERMINAL LEVEL SUPPORTS CHECKS
//=======================================
func supportsTrueColor(colorTerm string) bool {
	return colorTerm == "truecolor" || colorTerm == "24bit"
}

func supports256Color(term string) bool {
	return strings.Contains(term, "256color")
}

func supports16Color(term string) bool {
	return strings.Contains(term, "16color") || supportsColor(term)
}

func supportsColor(term string) bool {
	colorTerm := []string{"screen", "xterm", "vt100", "color", "ansi", "cygwin", "linux"}
	for _, t := range colorTerm {
		if strings.Contains(term, t) {
			return true
		}
	}
	return false
}
 
func supportsNone(term string) bool {
	return term == "dumb" || strings.Contains(term, "mono")
}

//Helper
func envExists(key string) bool {
	_, exists := os.LookupEnv(key)
	return exists
}


//WIP
//func windowsColorSupport() bool {
//	_, exists := os.LookupEnv("ANSICON")
//	return exists
//}


