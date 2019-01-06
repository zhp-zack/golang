package main

import (
	"fmt"
)

func test(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return test(n-1) + test(n-2)
	}
}

func main() {
	n := 4
	v := test(n)
	fmt.Println("v = ", v)
}
