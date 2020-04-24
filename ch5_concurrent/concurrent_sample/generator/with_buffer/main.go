package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func main() {
	ch := GenerateIntA()
	for i := 0; i < 10; i++ {
		fmt.Println(i, <-ch)
	}
}
