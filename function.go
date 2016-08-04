package main

import (

)

// We can specify multiple variable
// Return type is specified at the end
func multiply(a, b int) int {
	return a * b
}

func main() {
	product := multiply(10, 20)
	println(product)
}