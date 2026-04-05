package crayon

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"os"
	//uncomment after moving to version 1.24
	"golang.org/x/term"
)

type TempPart struct {
  Text string
  Index int
}

type CompiledTemplate struct {
  Parts []TempPart
  TotalLength int
}

type ColorToggle struct {
  EnableColor bool
}

func autoDetect() bool {
  if _, exists := os.LookupEnv("NO_COLOR"); exists{
    return false	
  }
  //uncomment after moving to version 1.24
  return term.isTerminal(int(os.Stdout.Fd()))//{

  //}
  //comment after moving to version 1.24
  //fileInfo, _ := os.Stdout.Stat()
  //os.ModeCharDevice should be replace with a cross platform version because this behaves differently on windows.
  //return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

//should auto detect tty by default
func NewColorToggle(enableColor ...bool) *ColorToggle {
  var colorEnabled bool
  if len(enableColor) > 0{
	colorEnabled = enableColor[0]
  } else {
	colorEnabled = autoDetect()
  }
  return &ColorToggle{
	EnableColor: colorEnabled,
  }
}


func (toggle *ColorToggle) Parse(input string) CompiledTemplate {
  if toggle == nil {
	toggle = NewColorToggle()
  }
  
  var (
	contentSequence  = ""
	inReadSequence   = false
	parts            []TempPart
	currentText      = ""
	allWords          []string
  )

  for i, ch := range input {
	char := string(ch)
	if char == "[" && !inReadSequence{
	  //check if the next value is "["
      // [[fg=color]] should never be an escape
      //consider first '[' as a text, move until, content is found. 
	  if i+1 < len(input) && input[i+1] == '['{
		currentText += "["
		continue
	  } else {
		inReadSequence = true
		contentSequence = ""
		allWords = nil

		if len(currentText) > 0 {
		  parts = append(parts, TempPart{Text: currentText, Index: -1})
		  currentText = ""
		}
	  }
	} else if ch == ']' && inReadSequence {
	    inReadSequence = false
		//if last word is present, add it
		allWords = strings.Fields(contentSequence)

		//check if all in [] are colors
		allColors := len(allWords) > 0
		for _, w := range allWords{
		  if !IsSupportedColor(w){
			allColors = false
			//break
		  }
		}
		if allColors{
		  if toggle.EnableColor {
			for _, w := range allWords{
			  parts = append(parts, TempPart{Text: ParseColor(w), Index: -1})
			}
		  } else {
			//redirected output or force turn off color
			parts = append(parts, TempPart{Text: "", Index: -1})
		  }
		} else {
			//not a color
		  if len(contentSequence) > 0 && allDigits(contentSequence){
			//decided to make it flexible and accept more indices but its still prone to overflow
            //needs a digit boundary guard	
			index, err := strconv.Atoi(contentSequence)
			// limit for indices
			if err == nil && index >= 0 && index <= 999 {
			  parts = append(parts, TempPart{Text: "", Index: index})
			} else {
			  addText := "[" + contentSequence + "]"
			  parts = append(parts, TempPart{Text: addText, Index: -1})
			}
		  } else{
			addText := "[" + contentSequence + "]"
			parts = append(parts, TempPart{Text: addText, Index: -1})
		  }
		}
	} else if inReadSequence {
	  contentSequence += char
	} else{
	  currentText += char
	}
  }

  if len(currentText) > 0 {
	parts = append(parts, TempPart{Text: currentText, Index: -1})
  }

  return CompiledTemplate{
	Parts: parts,
	TotalLength: len(input),
  }
}


//Override - without explicit toggle
func Parse(input string) CompiledTemplate {
  return NewColorToggle().Parse(input)
}
  

func allDigits(s string) bool {
  for _, r := range s{
	if !unicode.IsDigit(r){
	  return false
	}
  }
  return true
}
  

//apply will be a private func
func (temp CompiledTemplate) apply(args ...any) string {
  //Calculate estimated size for optimization
  var totalArgLength int
  for _, arg := range args{
	totalArgLength += len(fmt.Sprint(arg))
  }

  estimatedSize := temp.TotalLength + totalArgLength
  var result strings.Builder
  result.Grow(estimatedSize)

  for _, part := range temp.Parts{
	if part.Index < 0{
	  result.WriteString(part.Text)
	} else {
	  if part.Index < len(args) {
		result.WriteString(fmt.Sprint(args[part.Index]))
	  }
	}
  }
  return result.String()
}



//=======================
// PRINT
//=======================
func (temp CompiledTemplate) Println(args ..any) {
	fmt.Println(temp.apply(args...))
}

func (temp CompiledTemplate) Printf(format string, args ..any) {
	fmt.Printf(format, temp.apply(args...))
}

func (temp CompiledTemplate) Print(args ..any) {
	fmt.Print(temp.apply(args...))
}




//=======================
// EPRINT
//=======================
func (temp CompiledTemplate) Eprintln(args ..any) {
	fmt.Fprintln(os.Stderr, temp.apply(args...))
}

func (temp CompiledTemplate) Eprintf(format string, args ..any) {
	fmt.Fprintf(os.Stderr, format, temp.apply(args...))
}

func (temp CompiledTemplate) Eprint(args ..any) {
	fmt.Fprint(os.Stderr, temp.apply(args...))
}



//=======================
// FPRINT
//=======================
func (temp CompiledTemplate) Fprintln(w io.Writer, args ..any) (n int, err error) {
	return fmt.Fprintln(w, temp.apply(args...))
}

func (temp CompiledTemplate) Fprintf(w io.Writer, format string, args ..any) (n int, err error) {
	return fmt.Fprintf(w, format, temp.apply(args...))
}

func (temp CompiledTemplate) Fprint(w io.Writer, args ..any) (n int, err error){
	return fmt.Fprint(w, temp.apply(args...))
}


//=======================
// SPRINT
//=======================
func (temp CompiledTemplate) Sprintln(args ..any) string {
	return fmt.Sprintln(temp.apply(args...))
}

func (temp CompiledTemplate) Sprintf(args ..any) string{
	return fmt.Sprintln(temp.apply(args...))
}

func (temp CompiledTemplate) Sprint(args ..any) string {
	return temp.apply(args...)
}