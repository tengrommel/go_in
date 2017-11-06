package main

import (
	"fmt"
)

type power struct {
	attack int
	defense int
}

type location struct {
	x float32
	y float32
	z float32
}

type nonPlayerCharacter struct {
	name string
	speed int
	hp int
	power power
	loc location
}

func main() {
	fmt.Println("Structs...")
	demo := &nonPlayerCharacter{
		name: "草莓",
		speed: 21,
		hp: 1000,
		power: power{attack:75, defense: 50},
		loc: location{x: 1075.123, y: 521.123, z: 211.231},
	}
	fmt.Printf("实例的名字：%v\n 实例的hp值%v",demo.name, demo.hp)


	fmt.Println("Structs...")
	nothing_demo := &nonPlayerCharacter{
		name: "周腾",
		speed: 28,
		hp: 20000,
		power: power{attack:15, defense: 80},
		loc: location{x: 1074.123, y: 521.126, z: 212.231},
	}
	fmt.Printf("实例的名字：%v\n 实例的hp值%v",nothing_demo.name, nothing_demo.hp)
}
