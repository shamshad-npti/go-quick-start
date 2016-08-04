package main

import (
	"fmt"
)

func main() {
	// Create an array and initialize it
	// Array are indexed from 0
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	
	// Create an int array of size 5
	mums := [5]int{}
	mums[0] = 20
	fmt.Println(mums)

	// Create an empty array
	var sums []int
	fmt.Println(sums)

	// Append a value at the end of the array
	sums = append(sums, 20)
	fmt.Println(sums)

	// Create an array and assign first element later on
	var rums [5]int
	rums[0] = 11
	fmt.Println(rums)

	// Create an array of size 0 and capacity
	dums := make([]int, 0, 10)
	dums = append(dums, 10)
	fmt.Println(dums)

	// Create an array of size 10 and set 8th element
	gums := make([]int, 10)
	gums[8] = 10
	fmt.Println(gums)

	// Create an array of 8 element
	aums := [10]int {1, 2, 3, 4, 5, 6, 7, 8}
	slice := aums[0:5]
	slice = append(slice, 10)
	slice[0] = -1
	// Slice is modifed as expected
	fmt.Println(slice)
	// Aums is also modified [Unexpected for python programmer]
	fmt.Println(aums)
}