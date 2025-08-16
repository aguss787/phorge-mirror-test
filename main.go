package main

import "fmt"

func hello(i int) string {
	return fmt.Sprintf("Hello %d", i)
}

func main() {
	fmt.Println(hello(0))
	fmt.Println("Hello world")
	fmt.Println(hello(2))
	fmt.Println(hello(4))
	fmt.Println(hello(6))
	fmt.Println(hello(8))
	fmt.Println(hello(10))
	fmt.Println(hello(12))
}
