package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA(done chan struct{}) chan int {
	var ch = make(chan int)
	go func() {
		Label:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	var done = make(chan struct{})
	var ch = GenerateIntA(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(done)
	fmt.Println(<-ch)
}
