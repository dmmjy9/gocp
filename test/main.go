package main

import "runtime"

func count(c chan struct{}, ci chan int) {
	for i := 0; i < 10; i++ {
		ci <- i
	}
	close(ci)
	c <- struct{}{}
}

func main() {
	c := make(chan struct{})
	ci := make(chan int, 10)
	go count(c, ci)
	println("NumGoroutine: ",runtime.NumGoroutine())

	for v := range ci {
		println(v)
	}
}
