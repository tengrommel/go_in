package main

import (
	"fmt"
)

func main() {
	var x, y int = 3, 4
	if z:= x*y; z > 10{
		fmt.Println(z)
	} else {
		fmt.Println("failed:", z)
	}
}
