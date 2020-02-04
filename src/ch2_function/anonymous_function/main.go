package main

import (
	"fmt"
)

func main() {
	t := func(a, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(t)
}
