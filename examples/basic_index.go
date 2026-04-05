package main

import (
    "github.com/ph4mished/crayon"
)

func main() {
    // Simple template with one placeholder
    greeting := crayon.Parse("[fg=green]Hello, [0][reset]!")
    
    greeting.Println("Alice")
    greeting.Println("Bob")
    greeting.Println("World")
    
    // Complex template with multiple placeholders
    logTemplate := crayon.Parse("[0] [fg=blue][1][reset]: [fg=yellow][2][reset]")
    
    // Different log levels
    logTemplate.Println("[INFO]", "main", "Application started")
    logTemplate.Println("[WARN]", "auth", "Token expiring soon")
    logTemplate.Println("[ERROR]", "db", "Connection failed")
}