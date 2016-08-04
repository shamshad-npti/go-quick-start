package main
import (
	"fmt"
)
type student struct { name string; grade float32 }

func compare_student(left, right *student) int {
	if left.grade == right.grade {
		return 0
	} else if left.grade > right.grade {
		return 1
	} else {
		return -1
	}
}

func re_order(a, b int) (int, int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

func main() {
	stud := student{"Shamshad", 4.5}
	amit := student{grade: 4.6}
	anil := student{name: "Anil"}
	fmt.Printf("Name: %s\nGrade: %f\n", stud.name, stud.grade)
	fmt.Printf("Name: %s\nGrade: %f\n", amit.name, amit.grade)
	fmt.Printf("Name: %s\nGrade: %f\n", anil.name, anil.grade)

	reff := &stud
	fmt.Printf("Name: %s\nGrade: %f\n", reff.name, reff.grade)

	println(compare_student(&stud, &amit))
	a, b := re_order(30, 20)
	println(a, b)
}