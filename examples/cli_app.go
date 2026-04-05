package main
import(
    "os"
    "github.com/ph4mished/crayon"
)
// Best practice for CLI applications
func main() {
    // Check for --no-color flag
    noColorFlag := false
    for _, arg := range os.Args {
        if arg == "--no-color" {
            noColorFlag = true
            break
        }
    }
    
    // Respect both flag and environment variable
    useColor := !noColorFlag && os.Getenv("NO_COLOR") == ""
    
    // Create toggle
    toggle := crayon.NewColorToggle(useColor)
    
    // All templates use this toggle
    templates := struct {
        Success crayon.CompiledTemplate
        Error   crayon.CompiledTemplate
        Header  crayon.CompiledTemplate
    }{
        Success: toggle.Parse("[fg=green]✓ [0][reset]"),
        Error:   toggle.Parse("[fg=red]✗ [0][reset]"),
        Header:  toggle.Parse("[bold][0][reset]"),
    }
    
    // Use templates - they'll respect the toggle
    templates.Header.Println("My Application")
    templates.Success.Println("Started successfully")
    
    // If --no-color was used or NO_COLOR is set,
    // outputs will be plain text without escape codes
}