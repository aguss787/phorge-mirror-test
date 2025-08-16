package main

import "fmt"

func hello(i int) string {
	return fmt.Sprintf("Hello %d", i)
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(hello(1))
	fmt.Println(hello(2))
	fmt.Println(hello(3))
}
