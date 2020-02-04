package main

import (
	"fmt"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	ret := 0
	for i := 0; i < n; i++ {
		ret = fib(n-2) + fib(n-1)
	}
	return ret
}

func main() {
	ret := fib(5)
	fmt.Println(ret)
}
