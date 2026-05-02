package inkstamp

import (
	//"bytes"
	//"strings"
	"testing"
	//"github.com/inkstamp/inkstamp/termcolor"
)

//var colorCap = termcolor.Capability()

//=================================
// PARSE LOOP TESTS
//=================================

func TestParseLoop_BasicText(t *testing.T) {
	parts, text := parseLoop("Hello World", true, colorCap)

	if len(parts) != 0 {
		t.Errorf("Expected 0 parts, got %d", len(parts))
	}

	if text != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", text)
	}
}

func TestParseLoop_SimpleColorTag(t *testing.T) {
	parts, text := parseLoop("[fg=red]Hello", true, colorCap)

	if len(parts) != 1 {
		t.Errorf("Expected 1 part, got %d", len(parts))
	}

	if text != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", text)
	}

	if parts[0].Index != -1 {
		t.Errorf("Expected index -1, got %d", parts[0].Index)
	}
}

func TestParseLoop_MultipleColorTags(t *testing.T) {
	parts, text := parseLoop("[fg=red][bg=blue]Multi", true, colorCap)

	if len(parts) != 2 {
		t.Errorf("Expected 2 parts, got %d", len(parts))
	}

	if text != "Multi" {
		t.Errorf("Expected 'Multi', got '%s'", text)
	}
}


func TestParseLoop_ColorWithTextBeforeAndAfter(t *testing.T) {
	parts, text := parseLoop("Start [fg=cyan]Middle[reset] End", true, colorCap)

	if len(parts) != 4 {
		t.Errorf("Expected 4 parts, got %d", len(parts))
	}

	if text != " End" {
		t.Errorf("Expected ' End', got '%s'", text)
	}
}


//==================================
// BRACKET HANDLER TESTS
//==================================

func TestHandleOpenBrackets_DoubleBracket(t *testing.T) {
	parts := []TempPart{}
	currentText := ""

	parts, currentText, _, inSeq := handleOpenBracket(0, "[[fg=red]]", parts, currentText)

	if currentText != "[" {
		t.Errorf("Expectec currentText '[', got '%s'", currentText)
	}
	if inSeq {
		t.Errorf("Expected inSeq false for double brackets")
	}
}


func TestHandleOpenBrackets_ValidOpen(t *testing.T) {
	parts := []TempPart{}
	currentText := "[fg=red]"

	parts, currentText, _, inSeq := handleOpenBracket(0, "before [fg=red]", parts, currentText)

	if len(parts) != 1 {
		t.Errorf("Expected 1 part (flushed text), got %d", len(parts))
	}

	if currentText != "" {
		t.Errorf("Expected empty currentText, got '%s'", currentText)
	}
	if !inSeq {
		t.Errorf("Expected inSeq true")
	}
}


func TestHandleCloseBrackets_ColorSequence(t *testing.T) {
	parts := []TempPart{}
	parts, _ = handleCloseBracket("fg=red bg=blue", parts, true, colorCap)

	if len(parts) != 2 {
		t.Errorf("Expected 2 parts for two color tags, got %d", len(parts))
	}
}

func TestHandleCloseBrackets_ColorSequenceDisabled(t *testing.T) {
	parts := []TempPart{}
	parts, _ = handleCloseBracket("fg=red bg=blue", parts, false, colorCap)

	if len(parts) != 1 {
		t.Errorf("Expected 1 empty part, got %d", len(parts))
	}

	for i, part := range parts {
		if part.Text != ""{
			t.Errorf("Part %d: expected empty text, got '%s'", i, part.Text)
		}
	}
}




//=================================
// ALIGN TESTS
//=================================

func TestLeftAlign_WithFillChar(t *testing.T) {
	leftAligned := leftAlign("Hello World", 20, "=")
	expectedLen := 20

	if len(leftAligned) != expectedLen {
		t.Errorf("Expected length of leftAligned as %d, got %d", expectedLen, len(leftAligned))
	}

	if leftAligned != "Hello World=========" {
		t.Errorf("Expected 'Hello World========', got '%s'", leftAligned)
	}
}


func TestRightAlign_WithFillChar(t *testing.T) {
	rightAligned := rightAlign("Hello World", 20, "=")
	expectedLen := 20

	if len(rightAligned) != expectedLen {
		t.Errorf("Expected length of rightAligned as %d, got %d", expectedLen, len(rightAligned))
	}

	if rightAligned != "=========Hello World" {
		t.Errorf("Expected '========Hello World', got '%s'", rightAligned)
	}
}

func TestCenterAlign_WithFillChar(t *testing.T) {
	centerAligned := centerAlign("Hello, World", 20, "=")
	expectedLen := 20

	if len(centerAligned) != expectedLen {
		t.Errorf("Expected length of centerAligned as %d, got %d", expectedLen, len(centerAligned))
	}

	if centerAligned != "====Hello, World====" {
		t.Errorf("Expected '====Hello, World====', got '%s'", centerAligned)
	}
}



// =============================
// BENCHMARK TESTS
// =============================

func BenchmarkHandleCloseBrackets_ColorSequenceDisabled(b *testing.B) {
	parts := []TempPart{}
	for i := 0; i < b.N; i++ {
		_, _ = handleCloseBracket("fg=red bg=blue", parts, false, colorCap)
	}
}

func BenchmarkHandleCloseBrackets_ColorSequence(b *testing.B) {
	parts := []TempPart{}
	for i := 0; i < b.N; i++ {
		_, _ = handleCloseBracket("fg=red bg=blue", parts, true, colorCap)
	}
}

func BenchmarkHandleOpenBrackets_ValidOpen(b *testing.B) {
	parts := []TempPart{}
	currentText := "[fg=red]"

	
	for i := 0; i < b.N; i++ {
		_, _, _, _ = handleOpenBracket(0, "before [fg=red]", parts, currentText)
	}
}







func BenchmarkParseLoop_ColorWithTextBeforeAndAfter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = parseLoop("Start [fg=cyan]Middle[reset] End", true, colorCap)
	}
}

func BenchmarkParseLoop_MultipleColorTags(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = parseLoop("[fg=red][bg=blue]Multi", true, colorCap)
	}
}

func BenchmarkParseLoop_SimpleColorTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = parseLoop("[fg=red]Hello", true, colorCap)
	}
}



func BenchmarkApply(b *testing.B) {
	var crayonRed = Parse("[fg=red][0][reset]")
	for i := 0; i < b.N; i++ {
		crayonRed.apply("Hello world!")
	}
}

func BenchmarkSprint(b *testing.B) {
	var crayonRed = Parse("[fg=red][0][reset]")
	for i := 0; i < b.N; i++ {
		crayonRed.Sprint("Hello world!")
	}
}







