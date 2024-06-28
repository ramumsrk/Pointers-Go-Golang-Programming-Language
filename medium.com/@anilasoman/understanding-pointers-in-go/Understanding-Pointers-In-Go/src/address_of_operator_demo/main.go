package main

import (
  "fmt"
)

/*
 * Go program execution starts
 * at `main()' function
 */
func main() {
    var natural_number int = 32
    fmt.Println("Value of identifier 'natural_number' is", natural_number)
    fmt.Println("Address of identifier 'natural_number' is", &natural_number)
  return
}
