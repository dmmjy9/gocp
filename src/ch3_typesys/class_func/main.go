package main

import (
	"fmt"
)

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(i int) {
	t.a = i
}

func main() {
	var t = &T{}
	t.Set(2)
	fmt.Println(t.Get())
}
