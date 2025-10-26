package hello

import (
	"fmt"
)

func privateTest() {
	fmt.Println("Hello Test!")
}

func SayTest() {
	privateTest()
}
