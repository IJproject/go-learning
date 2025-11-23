package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func main() {
	Q1()
	Q2()
}

func (v Vertex) Plus() int {
	return v.X + v.Y
}
func Q1() {
	v := Vertex{X: 3, Y: 4}
	fmt.Println(v.Plus()) // => 7
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %d! Y is %d!", v.X, v.Y)
}
func Q2() {
	v := Vertex{3, 4}
	fmt.Println(v) // => X is 3! Y is 4!
}
