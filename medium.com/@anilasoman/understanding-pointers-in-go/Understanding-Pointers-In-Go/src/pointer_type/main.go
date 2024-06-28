package main

import (
  "fmt"
  "reflect"
)

/*
 * Go program execution starts
 * from `main()' function
 */
func main() {
    natural_number := 10
    fmt.Println("Value of identifier 'natural_number' is", natural_number)
    fmt.Println("Address of identifier 'natural_number' is", &natural_number)
    fmt.Println("Data-type of identifier 'natural_number' is", reflect.TypeOf(&natural_number))
  return
}
