package main

import (
	//"fmt"
)

func main() {
	// Declare a variable
	var power float32
	// Assign a variable
	power = 0.8
	// declare and assign
	name := "Revenger"
	println("My power is power", power)
	println("My name is name", name)
	voltage := complex(0.5, 0.9)
	current := complex(0.6, 0.4)
	e_power := voltage * current
	println(voltage, current, real(e_power))
}