package termcolor

import (
	//"os"
	"testing"
)

//CHECK

//detectColorCap
//checkTermEnv
//getForceColorLevel
//getCIColorLevel
//getKnownTerminal
//checkTermProgram
//queryTput


//MINI FUNCTIONS
//supportsTrueColor
//supports256Color
//supports16Color
//supportsNone
//supportsColor
//envExists

func TestCheckTermEnv(t *testing.T){
tests := []struct {
		name string
		term string
		colorTerm string
		expectedCap ColorCap
		expectedBool bool
	}{
		{"truecolor", "", "truecolor", ColorTrue, true},
		{"24bit", "", "24bit", ColorTrue, true},
		{"256color", "xterm-256color", "", Color256, true},
		{"16color", "xterm", "", Color16, true},
		{"dumb", "dumb", "", ColorNone, true},
		{"unknown", "unknown", "", unknown, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
		t.Setenv("TERM", tt.term)
		t.Setenv("COLORTERM", tt.colorTerm)

		cap, ok := checkTermEnv()
		if cap != tt.expectedCap || ok != tt.expectedBool{
			t.Errorf("expected (%v, %v), got (%v, %v)", tt.expectedCap, tt.expectedBool, cap, ok)
		}
	})
	}
}