package main

import (
	"runtime"
	"time"
)

func sum() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
		println(sum)
	}
}

func main() {
	go sum()
	println("NumGoroutine =", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
}
