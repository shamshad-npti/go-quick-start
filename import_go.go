package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("Hello", os.Args[1], "\b!")
	fmt.Println("I am safe at", os.Args[0])
}