package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p Person) walk() {
	fmt.Println("WALK")
}

func (p Person) showName() {
	fmt.Println(p.name)
}

func (p Person) showAge() {
	fmt.Println(p.age)
}

func (p Person) nameAddOne() int {
	return p.age + 1
}

func main() {
	p := Person{"andes", 18}
	p.walk()
	p.showAge()
	p.showName()
	fmt.Println(p.nameAddOne())
}
