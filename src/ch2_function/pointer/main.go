package main

import (
	"fmt"
)

func add(t *int) {
	*t = *t + 1
}

func main() {
	t := 1
	add(&t)
	fmt.Println(t)
}
