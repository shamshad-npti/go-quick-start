package main 

import (
	"fmt"
)

type arith struct {a, b int}

func (a arith) Add() int {
	return a.a + a.b
}

func main() {
	s := arith{4, 3}
	fmt.Println(s.Add())
}