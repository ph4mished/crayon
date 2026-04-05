package main

import (
    "fmt"
    "github.com/ph4mished/crayon"
)

func main() {
    // Parse and use color codes directly
    red := crayon.ParseColor("fg=red")
    bold := crayon.ParseColor("bold")
    reset := crayon.ParseColor("reset")
    
    fmt.Printf("%sThis is red and bold!%s\n", red + bold, reset)

    // Check if a color is supported
    if crayon.IsSupportedColor("fg=#FF0000") {
        fmt.Println("Hex colors are supported!")
    }
    
    // Or use the main functions
    crayon.Parse("[fg=blue]Hello in blue![reset]").Println()
    crayon.Parse("[bg=yellow fg=black bold]Bold black text on yellow background.[reset]").Println()


    //Or pre-parse the color template with placeholders for reuse. This is the heart of the library's performance.

    // Parse once
    template := crayon.Parse("[fg=red bold]Error: [0][reset]")

    // Reuse multiple times
    template.Println("File not found")
    template.Println("Permission denied")
    template.Println("Network timeout")
}