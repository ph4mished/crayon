package termcolor

import (
	//"os"
	"testing"
)

//CHECK
//checkColorToggle
//matchNextArg
//getNextArg
//isNumeric


//=============================================
// TESTING GET_NEXT_ARG FUNCTION
//=============================================
func TestGetNextArg_GNU(t *testing.T) {
	cachedArgs := []string{"--format=json", "--color=true"}
	value, ok := getNextArg(cachedArgs, "color", "=")
	if !ok {
		t.Error("getNextArg() return true for GNU style flags")
	}
	if value == "" {
		t.Errorf("Expected a value from getNextArg(), got='%s'", value)
	}
}

func TestGetNextArg_POSIX(t *testing.T) {
	cachedArgs := []string{"--format", "json", "--color", "true"}
	value, ok := getNextArg(cachedArgs, "color", "=")
	if !ok {
		t.Error("getNextArg() should return true for POSIX style flags")
	}
	if value == "" {
		t.Errorf("Expected a value from getNextArg(), got='%s'", value)
	}
}

func TestGetNextArg_NoValue(t *testing.T) {
	cachedArgs := []string{"--format", "json", "--color"}
	value, ok := getNextArg(cachedArgs, "color", "=")
	if !ok {
		t.Error("getNextArg() should return true for flags without values which implies bool flags")
	}
	if value != "" {
		t.Errorf("Expected empty value from getNextArg(), got='%s'", value)
	}
}


func TestGetNextArg_Hyphen_POSIX(t *testing.T) {
	cachedArgs := []string{"--format", "json", "--color", "-"}
	value, ok := getNextArg(cachedArgs, "color", "=")
	if !ok {
		t.Error("getNextArg() should return true for hyphens '-' as they are values too")
	}
	if value == "" {
		t.Errorf("Expected a value from getNextArg(), got='%s'", value)
	}
}


func TestGetNextArg_Hyphen_GNU(t *testing.T) {
	cachedArgs := []string{"--format=json", "--color=-"}
	value, ok := getNextArg(cachedArgs, "color", "=")
	if !ok {
		t.Error("getNextArg() should return true for hyphens '-' as they are values too")
	}
	if value == "" {
		t.Errorf("Expected a value from getNextArg(), got='%s'", value)
	}
}


func TestGetNextArg_Negative_POSIX(t *testing.T) {
	cachedArgs := []string{"--format", "json", "--number", "-6"}
	value, ok := getNextArg(cachedArgs, "number", "=")
	if !ok {
		t.Error("getNextArg() should return true for negatives numbers as they are values too")
	}
	if value == "" {
		t.Errorf("Expected a value from getNextArg(), got='%s'", value)
	}
}


func TestGetNextArg_Negative_GNU(t *testing.T) {
	cachedArgs := []string{"--format=json", "--number=-6"}
	value, ok := getNextArg(cachedArgs, "number", "=")
	if !ok {
		t.Error("getNextArg() should return true for negatives numbers as they are values too")
	}
	if value == "" {
		t.Errorf("Expected a value from getNextArg(), got='%s'", value)
	}
}

func TestGetNextArg_NextFlag_POSIX(t *testing.T) {
	cachedArgs := []string{"--color", "--format", "json"}
	value, ok := getNextArg(cachedArgs, "color", "=")
	if !ok {
		t.Error("getNextArg() should return true if nextArg is a flag, implies currentArg is a bool flag")
	}
	if value != "" {
		t.Errorf("Expected empty value from getNextArg(), got='%s'", value)
	}
}

func TestGetNextArg_NextFlag_GNU(t *testing.T) {
	cachedArgs := []string{"--color=", "--format=json"}
	value, ok := getNextArg(cachedArgs, "color", "=")
	if !ok {
		t.Error("getNextArg() should return true if nextArg is a flag, implies currentArg is a bool flag")
	}
	if value != "" {
		t.Errorf("Expected empty value from getNextArg(), got='%s'", value)
	}
}


//==============================
// TESTING IS_NUMERIC FUNCTION
//==============================

func TestIsNumeric_String(t *testing.T) {
	ok := isNumeric("hello")
	if ok {
		t.Errorf("Expected false from isNumeric as given parameter is string.")
	}
}


func TestIsNumeric_PosInt(t *testing.T) {
	ok := isNumeric("9")
	if !ok {
		t.Errorf("Expected true from isNumeric as given parameter is int.")
	}
}

func TestIsNumeric_NegInt(t *testing.T) {
	ok := isNumeric("-9")
	if ok {
		t.Errorf("Expected false from isNumeric as given parameter is negative int.")
	}
}


//===================================
// TESTING MATCH_NEXT_ARG FUNCTION
//===================================

func TestMatchNextArg_Exists(t *testing.T) {
	seqArgs := []string{"json", "table", "default", "csv"}
	ok := matchNextArg("json", seqArgs)
	if !ok {
		t.Errorf("Expected matchNextArg() to return true as provided value exists in slice of strings")
	}
}


func TestMatchNextArg_NotExists(t *testing.T) {
	seqArgs := []string{"json", "table", "default", "csv"}
	ok := matchNextArg("yaml", seqArgs)
	if ok {
		t.Errorf("Expected matchNextArg() to return false as provided value does not exist in slice of strings")
	}
}


//===============================
// TESTING CHECK_COLOR_TOGGLE
//==============================
func TestCheckColorToggle_NoValue_POSIX(t *testing.T) {
	cachedArgs := []string{"--format", "json", "--color"}
	colorPosArgs = []string{"on", "true", "yes", "1", "always", "force"}
	//colorNegArgs = []string{"off", "false", "no", "0", "never", "none"}
	//colorAutoArgs = []string{"auto", "tty", "detect"}

	ok := checkColorToggle(cachedArgs, "color", "=", colorPosArgs)
	if !ok {
		t.Errorf("Expected checkColorToggle() to return true as --color bool is on")
	}
}

func TestCheckColorToggle_NoValue_GNU(t *testing.T) {
	cachedArgs := []string{"--format=json", "--color="}
	colorPosArgs = []string{"on", "true", "yes", "1", "always", "force"}

	ok := checkColorToggle(cachedArgs, "color", "=", colorPosArgs)
	if !ok {
		t.Errorf("Expected checkColorToggle() to return true as --color= bool is on")
	}
}

func TestCheckColorToggle_ValueMatch_POSIX(t *testing.T) {
	cachedArgs := []string{"--format", "json", "--color", "on"}
	colorPosArgs = []string{"on", "true", "yes", "1", "always", "force"}

	ok := checkColorToggle(cachedArgs, "color", "=", colorPosArgs)
	if !ok {
		t.Errorf("Expected checkColorToggle() to return true as --color value matches")
	}
}

func TestCheckColorToggle_ValueMatch_GNU(t *testing.T) {
	cachedArgs := []string{"--format=json", "--color=on"}
	colorPosArgs = []string{"on", "true", "yes", "1", "always", "force"}

	ok := checkColorToggle(cachedArgs, "color", "=", colorPosArgs)
	if !ok {
		t.Errorf("Expected checkColorToggle() to return true as --color value matches")
	}
}


func TestCheckColorToggle_ValueNonMatch_POSIX(t *testing.T) {
	cachedArgs := []string{"--format", "json", "--color", "never"}
	colorPosArgs = []string{"on", "true", "yes", "1", "always", "force"}

	ok := checkColorToggle(cachedArgs, "color", "=", colorPosArgs)
	if ok {
		t.Errorf("Expected checkColorToggle() to return false as --color value doesn't match")
	}
}


func TestCheckColorToggle_ValueNonMatch_GNU(t *testing.T) {
	cachedArgs := []string{"--format=json", "--color=never"}
	colorPosArgs = []string{"on", "true", "yes", "1", "always", "force"}

	ok := checkColorToggle(cachedArgs, "color", "=", colorPosArgs)
	if ok {
		t.Errorf("Expected checkColorToggle() to return false as --color value doesn't match")
	}
}
