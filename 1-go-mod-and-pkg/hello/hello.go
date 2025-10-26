package hello

import (
	"fmt"

	"go-mod/hello/internal/hi"
)

func SayHello() {
	fmt.Println("Hello!")
	hi.SayHi()
}
