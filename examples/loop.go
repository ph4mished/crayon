package main

import (
    "fmt"
    "time"
    "github.com/ph4mished/crayon"
)

func main() {
    const iterations = 1000000
    
    // Method 1: Parse once, apply many
    template := crayon.Parse("[bold fg=red][0][reset] [fg=green][1][reset]")
    
    start := time.Now()
    for i := 0; i < iterations; i++ {
        template.Sprint(fmt.Sprintf("Item%d", i), fmt.Sprintf("Value%d", i))
    }
    fmt.Printf("Template reuse: %v\n", time.Since(start))
    
    // Method 2: Parse every time
    start = time.Now()
    for i := 0; i < iterations; i++ {
        crayon.Parse(fmt.Sprintf("[bold fg=red]Item%d[reset] [fg=green]Value%d[reset]", i, i)).Sprint()
    }
    fmt.Printf("Parse every time: %v\n", time.Since(start))
    
    // Method 3: Manual concatenation
    start = time.Now()
    for i := 0; i < iterations; i++ {
        _ = crayon.ParseColor("fg=red bold") + fmt.Sprintf("Item%d", i) + 
            crayon.ParseColor("reset") + " " + 
            crayon.ParseColor("fg=green") + fmt.Sprintf("Value%d", i) + 
            crayon.ParseColor("reset")
    }
    fmt.Printf("Manual concatenation: %v\n", time.Since(start))
}