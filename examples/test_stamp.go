package main

import "github.com/phamio/crayon"
import "fmt"

func main(){
  // Simple colored text
  temp := crayon.Parse("[fg=blue]][Success] [fg=3]message [fg=#ffff00][copied][fg=rgb(255,255,0)][done][reset]")
  fmt.Println("MY DEFINITIONS: [fg=blue]][Success] [fg=3]message [fg=#ffff00][copied][fg=rgb(255,255,0)][done][reset]\n\n")
  fmt.Printf("%#v", temp)
  fmt.Printf("%+v", temp)
  fmt.Println("\n\n")
  }
  

