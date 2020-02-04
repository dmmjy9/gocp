package main

func main() {
	defer func() {
		println("first")
	}()

	defer func() {
		println("second")
	}()

	println("function body")
}