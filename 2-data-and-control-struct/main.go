package main

import "fmt"

func variable() {
	var name string = "Knot"
	var age int = 20

	var isMan = true

	// Shorthand
	weight := 80.1
	address := "Thailand"

	fmt.Printf("Hello %s age %d \n", name, age)
	fmt.Println("weight ", weight)
	fmt.Println("isMan ", isMan)
	fmt.Println("address ", address)

	address = "bangkok"
	fmt.Println("address ", address)

	// constant
	const pi = 3.14
	fmt.Println("pi ", pi)
}

func arithmeticOperator() {
	var a int = 10
	var b int = 20

	fmt.Println("a + b = ", a+b) // 30
	fmt.Println("a - b = ", a-b) // -10
	fmt.Println("a * b = ", a*b) // 200
	fmt.Println("a / b = ", a/b) // 0
	fmt.Println("a % b = ", a%b) // 10
}

func operator() {
	// Relational Operators
	var a int = 10
	var b int = 20

	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // false

	// Logical Operators
	fmt.Println(a == b && a > b) // false
	fmt.Println(a == b || a > b) // true
}

func controlStructure() {
	// if else
	var score int = 62
	var grade string

	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else {
		grade = "F"
	}

	fmt.Printf("Your grade is %s", grade)

	// switch case
	var dayOfWeek = 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}
}

func main() {
	variable()
}
