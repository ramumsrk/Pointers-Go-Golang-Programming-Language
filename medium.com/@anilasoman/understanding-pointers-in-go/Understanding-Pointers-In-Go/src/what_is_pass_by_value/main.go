package main

import "fmt"

/*
 * Go program execution
 * starts from `main()'
 * function
 */
func main() {
    var price float32 = 10
    fmt.Println("Value of 'price' before calling 'modifyAParameterPassedByValue'", price)
    // User-defined function call
    modifyAParameterPassedByValue(10)
    fmt.Println("Value of 'price' after calling 'modifyAParameterPassedByValue'", price)
  return
}

// User-defined function definition
func modifyAParameterPassedByValue(price float32) {
    price = 42
    fmt.Println("Value of 'price' inside 'modifyAParameterPassedByValue'", price)
  return
}
