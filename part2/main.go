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
}
