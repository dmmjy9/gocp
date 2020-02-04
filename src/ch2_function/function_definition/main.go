package main

import "fmt"

func add(arr ...int) (sum int) {
	for _, v := range arr{
		sum += v
	}
	return
}

func main() {
	i := []int{1, 3, 5, 7, 9}
	addRet := add(i...)
	fmt.Println(addRet)
}
