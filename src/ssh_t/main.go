package main

import "time"

func main() {
	go func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		time.Sleep(5 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}
