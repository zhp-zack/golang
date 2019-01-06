package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))
	fmt.Println("s = ", s)

	s1 := make([]int, 10, 10)
	fmt.Println("len(s1) = ", len(s1))
	fmt.Println("cap(s1) = ", cap(s1))
	fmt.Println("s1 = ", s1)

	var s2 []int
	var a [10]int
	s2 = a[:]
	fmt.Println("len(s2) = ", len(s2))
	fmt.Println("cap(s2) = ", cap(s2))
	fmt.Println("s2 = ", s2)

	s2 = append(s2, s[0])
	fmt.Println(s2)
}
