package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// 一个随机数生成器
func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			// 增加一路监听，就是对退出通知信号done的监听
			case <- done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 发送通知，告知生产者停止生产
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 生产者已退出
	fmt.Println("NumGoroutine = ",runtime.NumGoroutine())
}
