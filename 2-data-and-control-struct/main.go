package main

import (
	"fmt"
	"go-mod/methods"
)

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

	// pre-define if
	if num := 9; num < 0 {
		fmt.Println("Negative")
	} else if num < 10 {
		fmt.Println("Single digit")
	} else {
		fmt.Println("Multiple digits")
	}

	// iteration (Loop)
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// do while loop
	i = 0
	for {
		fmt.Println(i)
		i++
		if i >= 5 {
			break
		}
	}
}

func dataStructure() {
	// array
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)

	// slice (dynamic sequence)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = append(slice, 6)
	fmt.Println(slice)

	newSlice := arr[:]
	fmt.Println(newSlice)

	// map
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	fmt.Println(myMap)

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	// Check if key exist
	value, ok := m["one"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key not found")
	}

	// struct
	type Student struct {
		Name  string
		Age   int
		Grade string
	}
	var student Student
	student.Name = "John"
	student.Age = 20
	student.Grade = "A"
	fmt.Println(student)

	mapStudent := make(map[string]Student)
	mapStudent["st01"] = Student{Name: "Knot", Age: 20, Grade: "A"}
	fmt.Println(mapStudent)
}

func main() {
	methods.Receiver()
}
