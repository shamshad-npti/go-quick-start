package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := slice[3:6]
	fmt.Println(slice)
	fmt.Println(slice1)
	slice1[0] = 46
	fmt.Println(slice)
	slice1 = append(slice1, 20)
	fmt.Println(slice1)
	fmt.Println(slice)
	slice3 := make([]int, 5, 10)
}