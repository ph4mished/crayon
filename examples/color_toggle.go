package main

import (
    "os"
    "github.com/ph4mished/crayon"
)

func main() {
    // Create color toggle - respects NO_COLOR env var and when output is redirected by default
    toggle := crayon.NewColorToggle()
    
    // Parse templates using the toggle
    successTemplate := toggle.Parse("[fg=green]✓ [0][reset]")
    errorTemplate := toggle.Parse("[fg=red]✗ [0][reset]")
    
    // These will only show colors if appropriate
    successTemplate.Println("Operation completed")
    errorTemplate.Println("Operation failed")
    
    // Manual control
    //forceColors := crayon.NewColorToggle(true)   // Always show colors
    //noColors := crayon.NewColorToggle(false)     // Never show colors
    
    // Use in CLI applications
    useColor := os.Getenv("NO_COLOR") == ""
    appToggle := crayon.NewColorToggle(useColor)
    
    helpTemplate := appToggle.Parse("[bold fg=cyan][0][reset] [fg=green][1][reset]")
    helpTemplate.Println("Usage:", "myapp [options]")
}
