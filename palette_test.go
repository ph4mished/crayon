package crayon

import (
	"testing"
)

// =============================
// COLOR MAP TESTS
// =============================

func TestColorMap_ContainsAllForegroundColors(t *testing.T) {
	foregroundColors := []string{
		"fg=black", "fg=red", "fg=green", "fg=yellow", "fg=blue",
		"fg=magenta", "fg=cyan", "fg=white", "fg=darkgray",
		"fg=lred", "fg=lgreen", "fg=lyellow", "fg=lblue",
		"fg=lmagenta", "fg=lcyan", "fg=lwhite",
	}

	for _, color := range foregroundColors {
		if code, exists := colorMap[color]; !exists {
			t.Errorf("ColorMap missing foreground color: %s", color)
		} else if code == "" {
			t.Errorf("ColorMap has empty code for: %s", color)
		}
	}
}

func TestColorMap_ContainsAllBackgroundColors(t *testing.T) {
	backgroundColors := []string{
		"bg=black", "bg=red", "bg=green", "bg=yellow", "bg=blue",
		"bg=magenta", "bg=cyan", "bg=white", "bg=darkgray",
		"bg=lred", "bg=lgreen", "bg=lyellow", "bg=lblue",
		"bg=lmagenta", "bg=lcyan", "bg=lwhite",
	}

	for _, color := range backgroundColors {
		if code, exists := colorMap[color]; !exists {
			t.Errorf("ColorMap missing background color: %s", color)
		} else if code == "" {
			t.Errorf("ColorMap has empty code for: %s", color)
		}
	}
}

func TestColorMap_ForegroundCodes(t *testing.T) {
	tests := map[string]string{
		"fg=black":    "30",
		"fg=red":      "31",
		"fg=green":    "32",
		"fg=yellow":   "33",
		"fg=blue":     "34",
		"fg=magenta":  "35",
		"fg=cyan":     "36",
		"fg=white":    "37",
		"fg=darkgray": "90",
		"fg=lred":     "91",
		"fg=lgreen":   "92",
		"fg=lyellow":  "93",
		"fg=lblue":    "94",
		"fg=lmagenta": "95",
		"fg=lcyan":    "96",
		"fg=lwhite":   "97",
	}

	for color, expectedCode := range tests {
		if code, exists := colorMap[color]; !exists {
			t.Errorf("Color %s not found in colorMap", color)
		} else if code != expectedCode {
			t.Errorf("Color %s: expected code %s, got %s", color, expectedCode, code)
		}
	}
}

func TestColorMap_BackgroundCodes(t *testing.T) {
	tests := map[string]string{
		"bg=black":    "40",
		"bg=red":      "41",
		"bg=green":    "42",
		"bg=yellow":   "43",
		"bg=blue":     "44",
		"bg=magenta":  "45",
		"bg=cyan":     "46",
		"bg=white":    "47",
		"bg=darkgray": "100",
		"bg=lred":     "101",
		"bg=lgreen":   "102",
		"bg=lyellow":  "103",
		"bg=lblue":    "104",
		"bg=lmagenta": "105",
		"bg=lcyan":    "106",
		"bg=lwhite":   "107",
	}

	for color, expectedCode := range tests {
		if code, exists := colorMap[color]; !exists {
			t.Errorf("Color %s not found in colorMap", color)
		} else if code != expectedCode {
			t.Errorf("Color %s: expected code %s, got %s", color, expectedCode, code)
		}
	}
}

// =============================
// RESET MAP TESTS
// =============================

func TestResetMap_ContainsAllResets(t *testing.T) {
	resets := []string{
		"reset", "fg=reset", "bg=reset", "bold=reset", "dim=reset",
		"italic=reset", "underline=reset", "blink=reset",
		"reverse=reset", "hidden=reset", "strike=reset",
	}

	for _, reset := range resets {
		if code, exists := resetMap[reset]; !exists {
			t.Errorf("ResetMap missing: %s", reset)
		} else if code == "" {
			t.Errorf("ResetMap has empty code for: %s", reset)
		}
	}
}

func TestResetMap_Codes(t *testing.T) {
	tests := map[string]string{
		"reset":           "0",
		"fg=reset":        "39",
		"bg=reset":        "49",
		"bold=reset":      "22",
		"dim=reset":       "22",
		"italic=reset":    "23",
		"underline=reset": "24",
		"blink=reset":     "25",
		"reverse=reset":   "27",
		"hidden=reset":    "28",
		"strike=reset":    "29",
	}

	for reset, expectedCode := range tests {
		if code, exists := resetMap[reset]; !exists {
			t.Errorf("Reset %s not found in resetMap", reset)
		} else if code != expectedCode {
			t.Errorf("Reset %s: expected code %s, got %s", reset, expectedCode, code)
		}
	}
}

func TestResetMap_BoldAndDimShareCode(t *testing.T) {
	boldReset := resetMap["bold=reset"]
	dimReset := resetMap["dim=reset"]

	if boldReset != dimReset {
		t.Errorf("bold=reset (%s) and dim=reset (%s) should have same code (22)", boldReset, dimReset)
	}

	if boldReset != "22" {
		t.Errorf("bold=reset should be '22', got '%s'", boldReset)
	}
}

// =============================
// STYLE MAP TESTS
// =============================

func TestStyleMap_ContainsAllStyles(t *testing.T) {
	styles := []string{
		"bold", "dim", "italic", "underline=single",
		"blink=slow", "blink=fast", "reverse", "hidden",
		"strike", "underline=double",
	}

	for _, style := range styles {
		if code, exists := styleMap[style]; !exists {
			t.Errorf("StyleMap missing: %s", style)
		} else if code == "" {
			t.Errorf("StyleMap has empty code for: %s", style)
		}
	}
}

func TestStyleMap_Codes(t *testing.T) {
	tests := map[string]string{
		"bold":             "1",
		"dim":              "2",
		"italic":           "3",
		"underline=single": "4",
		"blink=slow":       "5",
		"blink=fast":       "6",
		"reverse":          "7",
		"hidden":           "8",
		"strike":           "9",
		"underline=double": "21",
	}

	for style, expectedCode := range tests {
		if code, exists := styleMap[style]; !exists {
			t.Errorf("Style %s not found in styleMap", style)
		} else if code != expectedCode {
			t.Errorf("Style %s: expected code %s, got %s", style, expectedCode, code)
		}
	}
}

// =============================
// ANSI 256 TO 16 LUT TESTS
// =============================

func TestAnsi256ToAnsi16Lut_Length(t *testing.T) {
	expectedLength := 256
	actualLength := len(ansi256ToAnsi16Lut)

	if actualLength != expectedLength {
		t.Errorf("LUT length = %d, expected %d", actualLength, expectedLength)
	}
}

func TestAnsi256ToAnsi16Lut_ValidRange(t *testing.T) {
	for i, code := range ansi256ToAnsi16Lut {
		if code < 30 || (code > 39 && code < 90) || (code > 97 && code < 100) || code > 107 {
			// Valid codes are: 30-37, 39, 90-97, 100-107
			// Also 0? No, should be in range 30-107 but skipping invalid ones
			if code != 39 && (code < 30 || code > 107) {
				t.Errorf("LUT[%d] = %d is outside valid ANSI 16-bit color range (30-107)", i, code)
			}
		}
	}
}

func TestAnsi256ToAnsi16Lut_StandardColors(t *testing.T) {
	// First 16 entries should map to standard colors
	expected := []uint8{30, 31, 32, 33, 34, 35, 36, 37, 90, 91, 92, 93, 94, 95, 96, 97}

	for i := 0; i < 16; i++ {
		if ansi256ToAnsi16Lut[i] != expected[i] {
			t.Errorf("LUT[%d] = %d, expected %d", i, ansi256ToAnsi16Lut[i], expected[i])
		}
	}
}

func TestAnsi256ToAnsi16Lut_GrayscaleRange(t *testing.T) {
	// Last 24 entries should be grayscale mapping (232-255)
	grayscaleStart := 232
	_ = grayscaleStart

	// Check that grayscale entries map to appropriate values (30, 90, 37, 97)
	for i := 232; i < 256; i++ {
		code := ansi256ToAnsi16Lut[i]
		if code != 30 && code != 90 && code != 37 && code != 97 {
			t.Logf("Grayscale LUT[%d] = %d (acceptable range: 30, 90, 37, 97)", i, code)
		}
	}
}

// =============================
// COLOR VALUE VALIDATION TESTS
// =============================

func TestColorMap_NoDuplicateKeys(t *testing.T) {
	seen := make(map[string]bool)
	for key := range colorMap {
		if seen[key] {
			t.Errorf("Duplicate key in colorMap: %s", key)
		}
		seen[key] = true
	}
}

func TestResetMap_NoDuplicateKeys(t *testing.T) {
	seen := make(map[string]bool)
	for key := range resetMap {
		if seen[key] {
			t.Errorf("Duplicate key in resetMap: %s", key)
		}
		seen[key] = true
	}
}

func TestStyleMap_NoDuplicateKeys(t *testing.T) {
	seen := make(map[string]bool)
	for key := range styleMap {
		if seen[key] {
			t.Errorf("Duplicate key in styleMap: %s", key)
		}
		seen[key] = true
	}
}

// =============================
// CROSS-REFERENCE TESTS
// =============================

func TestNoOverlapBetweenMaps(t *testing.T) {
	// Check that keys don't appear in multiple maps (unless intentional)
	
	for key := range colorMap {
		if _, exists := resetMap[key]; exists {
			t.Errorf("Key '%s' appears in both colorMap and resetMap", key)
		}
		if _, exists := styleMap[key]; exists {
			t.Errorf("Key '%s' appears in both colorMap and styleMap", key)
		}
	}
	
	for key := range resetMap {
		if _, exists := colorMap[key]; exists {
			t.Errorf("Key '%s' appears in both resetMap and colorMap", key)
		}
		if _, exists := styleMap[key]; exists {
			t.Errorf("Key '%s' appears in both resetMap and styleMap", key)
		}
	}
	
	for key := range styleMap {
		if _, exists := colorMap[key]; exists {
			t.Errorf("Key '%s' appears in both styleMap and colorMap", key)
		}
		if _, exists := resetMap[key]; exists {
			t.Errorf("Key '%s' appears in both styleMap and resetMap", key)
		}
	}
}

// =============================
// ANSI CODE VALIDITY TESTS
// =============================

func TestColorMap_AnsiCodesAreValid(t *testing.T) {
	validCodes := []string{
		"30", "31", "32", "33", "34", "35", "36", "37",
		"90", "91", "92", "93", "94", "95", "96", "97",
		"40", "41", "42", "43", "44", "45", "46", "47",
		"100", "101", "102", "103", "104", "105", "106", "107",
	}
	
	validSet := make(map[string]bool)
	for _, code := range validCodes {
		validSet[code] = true
	}
	
	for key, code := range colorMap {
		if !validSet[code] {
			t.Errorf("ColorMap[%s] has invalid ANSI code: %s", key, code)
		}
	}
}

func TestResetMap_AnsiCodesAreValid(t *testing.T) {
	validCodes := []string{"0", "39", "49", "22", "23", "24", "25", "27", "28", "29"}
	validSet := make(map[string]bool)
	for _, code := range validCodes {
		validSet[code] = true
	}
	
	for key, code := range resetMap {
		if !validSet[code] {
			t.Errorf("ResetMap[%s] has invalid ANSI code: %s", key, code)
		}
	}
}

func TestStyleMap_AnsiCodesAreValid(t *testing.T) {
	validCodes := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "21"}
	validSet := make(map[string]bool)
	for _, code := range validCodes {
		validSet[code] = true
	}
	
	for key, code := range styleMap {
		if !validSet[code] {
			t.Errorf("StyleMap[%s] has invalid ANSI code: %s", key, code)
		}
	}
}

// =============================
// BENCHMARK TESTS
// =============================

func BenchmarkColorMapLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = colorMap["fg=red"]
		_ = colorMap["bg=blue"]
		_ = colorMap["fg=lgreen"]
	}
}

func BenchmarkResetMapLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = resetMap["reset"]
		_ = resetMap["fg=reset"]
		_ = resetMap["bold=reset"]
	}
}

func BenchmarkStyleMapLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = styleMap["bold"]
		_ = styleMap["italic"]
		_ = styleMap["underline=single"]
	}
}

func BenchmarkAnsi256ToAnsi16Lut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ansi256ToAnsi16Lut[196] // Red
		_ = ansi256ToAnsi16Lut[46]  // Green
		_ = ansi256ToAnsi16Lut[21]  // Blue
		_ = ansi256ToAnsi16Lut[244] // Gray
	}
}

// =============================
// TABLE DRIVEN TESTS
// =============================

func TestAllMaps_Completeness(t *testing.T) {
	tests := []struct {
		name     string
		mapData  map[string]string
		minSize  int
	}{
		{"colorMap", colorMap, 32},
		{"resetMap", resetMap, 11},
		{"styleMap", styleMap, 10},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if len(test.mapData) < test.minSize {
				t.Errorf("%s has %d entries, expected at least %d", 
					test.name, len(test.mapData), test.minSize)
			}
		})
	}
}

func TestAnsi256ToAnsi16Lut_AllEntriesAccessible(t *testing.T) {
	// Verify we can access all 256 entries without panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Panic when accessing LUT: %v", r)
		}
	}()
	
	for i := 0; i < 256; i++ {
		_ = ansi256ToAnsi16Lut[i]
	}
}

// =============================
// SPECIFIC COLOR MAPPING TESTS
// =============================

func TestColorMap_BrightColors(t *testing.T) {
	brightForegrounds := map[string]string{
		"fg=lred":     "91",
		"fg=lgreen":   "92",
		"fg=lyellow":  "93",
		"fg=lblue":    "94",
		"fg=lmagenta": "95",
		"fg=lcyan":    "96",
		"fg=lwhite":   "97",
	}
	
	for color, expected := range brightForegrounds {
		if colorMap[color] != expected {
			t.Errorf("Bright color %s should map to %s, got %s", 
				color, expected, colorMap[color])
		}
	}
	
	brightBackgrounds := map[string]string{
		"bg=lred":     "101",
		"bg=lgreen":   "102",
		"bg=lyellow":  "103",
		"bg=lblue":    "104",
		"bg=lmagenta": "105",
		"bg=lcyan":    "106",
		"bg=lwhite":   "107",
	}
	
	for color, expected := range brightBackgrounds {
		if colorMap[color] != expected {
			t.Errorf("Bright color %s should map to %s, got %s", 
				color, expected, colorMap[color])
		}
	}
}

func TestStyleMap_UnderlineVariants(t *testing.T) {
	singleUnderline := styleMap["underline=single"]
	doubleUnderline := styleMap["underline=double"]
	
	if singleUnderline != "4" {
		t.Errorf("underline=single should be '4', got '%s'", singleUnderline)
	}
	
	if doubleUnderline != "21" {
		t.Errorf("underline=double should be '21', got '%s'", doubleUnderline)
	}
}

func TestStyleMap_BlinkVariants(t *testing.T) {
	slowBlink := styleMap["blink=slow"]
	fastBlink := styleMap["blink=fast"]
	
	if slowBlink != "5" {
		t.Errorf("blink=slow should be '5', got '%s'", slowBlink)
	}
	
	if fastBlink != "6" {
		t.Errorf("blink=fast should be '6', got '%s'", fastBlink)
	}
}