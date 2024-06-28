package main

import "fmt"

/*
 * Go program execution
 * starts from `main()'
 * function
 */
func main() {
    var price float32 = 10
    fmt.Println("Value of identifier 'price' before calling user-defined function 'modifyAParameterPassedAsAPointer'", price)
    modifyAParameterPassedAsAPointer(&price)
    fmt.Println("Value of identifier 'price' after calling user-defined function 'modifyAParameterPassedAsAPointer'", price)
  return
}

// User-defined function definition
func modifyAParameterPassedAsAPointer(price *float32) {
    *price = 42
    fmt.Println("Value of identifier 'price' in user-defined function 'modifyAParameterPassedAsAPointer'", *price)
  return
}
