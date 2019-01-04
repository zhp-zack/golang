package main

import (
	"fmt"
)

func main() {
	str := "hellow"
	fmt.Println("strlen = ", len(str))

	var p *int = new(int)
	fmt.Println("p = ", p)
}
