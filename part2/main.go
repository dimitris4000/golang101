package main

import (
	"fmt"
)

func main() {
	type Coord struct {
		X float32
		Y float32
	}

	c1 := Coord{X: 1, Y: 2}
	c2 := &c1              // <<< Get the memory address to c1 and put it in pc1
	fmt.Printf("%T\n", c2) // *main.Coord

	fmt.Println(c1) // {1 2}
	fmt.Println(c2) // {1 2}
	c1.X = 5
	fmt.Println(c1) // {5 2}
	fmt.Println(c2) // {5 2}

	add := func(a, b Coord) Coord {
		return Coord{a.X + b.X, a.Y + b.Y}
	}
	increase := func(a *Coord, b Coord) { // <<< Use pointer as a parameter
		a.X += b.X
		a.Y += b.Y
	}

	c3 := Coord{10, 20}
	res1 := add(c1, c3)
	c1.X++
	fmt.Println(c1)
	fmt.Println(res1)

	increase(&c1, c3) // <<< Pass a pointer parameter
	fmt.Println(c1)

	pc3 := &c3
	increase(&c1, *pc3) // <<< Pass a pointer as a value parameter (pc3)
	fmt.Println(c1)
}
