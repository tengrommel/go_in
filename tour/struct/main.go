package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X:1, Y:2}
	p := &v
	p.X = 999
	fmt.Println(v.X)
}
