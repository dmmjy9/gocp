package main

import "fmt"

type Person struct {
		name string
		age int
}

func main() {
		a := Person{"andes", 18}
		fmt.Println(a.name)
		fmt.Println(a.age)
}

