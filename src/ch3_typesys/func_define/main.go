package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

type ADD func(int, int) int

func main() {
	var g ADD = add
	a, b := 1, 3
	fmt.Println(g(a, b))
}
