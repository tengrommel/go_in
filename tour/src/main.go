package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	defer fmt.Println("world123")
	fmt.Println("hello1")
	defer fmt.Println("world456")
	defer fmt.Println("hello2")
}
