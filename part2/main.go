package main

import (
	"fmt"
)

type Coord struct {
	X float32
	Y float32
}

func NewCoord(x, y float32) Coord {
	return Coord{x, y}
}

func (c Coord) add(b Coord) Coord { // <<< Use value as a receiver
	return Coord{c.X + b.X, c.Y + b.Y}
}
func (c *Coord) increase(b Coord) { // <<< Use pointer as a receiver
	c.X += b.X
	c.Y += b.Y
}

func main() {
	c1 := NewCoord(1, 2)
	c2 := &c1
	fmt.Printf("%T\n", c2) // *main.Coord

	fmt.Println(c1) // {1 2}
	fmt.Println(c2) // {1 2}
	c1.X = 5
	fmt.Println(c1) // {5 2}
	fmt.Println(c2) // {5 2}

	c3 := NewCoord(10, 20)
	res1 := c1.add(c3)
	c1.X++
	fmt.Println(c1)
	fmt.Println(res1)

	c1.increase(c3)
	fmt.Println(c1)

	pc3 := &c3
	c1.increase(*pc3)
	fmt.Println(c1)
}
