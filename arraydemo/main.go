package main

import (
	"fmt"
)

func test(a *[3]int) {
	a[0] = 31
}

func main() {
	a := [3]int{1, 2, 3}
	test(&a)
	fmt.Println(a)
}
