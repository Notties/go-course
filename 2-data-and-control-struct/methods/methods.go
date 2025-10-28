package methods

import (
	"fmt"
)

// Define the Student struct
type Student struct {
	Firstname string
	Lastname  string
}

// Method with a receiver of type Student
// This method returns the full name of the student
func (s Student) FullName() string {
	return s.Firstname + " " + s.Lastname
}

func Receiver() {
	student := Student{
		Firstname: "Mike",
		Lastname:  "Lopster",
	}

	// Call the FullName method on the Student instance
	fullName := student.FullName()
	fmt.Println("Full Name of the student:", fullName)
}
