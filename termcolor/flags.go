package termcolor

import (
	"strings"
	"unicode"
)

var (
	colorPosArgs = []string{"on", "true", "yes", "1", "always", "force"}
	colorNegArgs = []string{"off", "false", "no", "0", "never", "none"}
	colorAutoArgs = []string{"auto", "tty", "detect"}
)


func checkColorToggle(cachedArgs []string, flag, delimiter string, seqArgs []string) bool {
	nextArg, boolean := getNextArg(cachedArgs, flag, delimiter)
	if boolean{
		if nextArg != ""{
		return matchNextArg(nextArg, seqArgs)
		}
		return true
	}
	return false
}



func matchNextArg(arg string, seqArgs []string) bool {
	for _, value := range seqArgs {
		if arg == value {
			return true
		}
	}
	return false
}



func getNextArg(cachedArgs []string, flag string, delimiter string) (string, bool) {
	flagDelim := flag+delimiter
	args := cachedArgs
	for i, arg := range args{
		cleanArg := strings.TrimLeft(arg, "-")
		//GNU Style
		if strings.HasPrefix(cleanArg, flagDelim){
			return strings.TrimPrefix(cleanArg, flagDelim), true
		}

        //POSIX Style
		if cleanArg == flag {
		    if i+1 < len(args) {
				nextArg := args[i+1]
				//Check if next arg is not a flag or is not a negative number or is not just a hyphen
				if !strings.HasPrefix(nextArg, "-") || isNumeric(nextArg[1:]) || nextArg == "-" {
					//Thats the value of given flag
					return nextArg, true
				}
			}
			//if flag has no value (bool)
			return "", true
		}
	}
	//Flag wasn't present
	return "", false
}


func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}