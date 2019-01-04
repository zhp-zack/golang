package main

import (
	"fmt"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err =", err)
		}
	}()

	n1 := 10
	n2 := 0
	n := n1 / n2
	fmt.Println("error handler", n)
}
