package main

import "fmt"

func main() {
	v3 := NewVertex3D(2, 3, 4)
	fmt.Println(v3.Area3D())   // => 24
}

type Vertex2D struct {
	X, Y int
}

// Embedded
type Vertex3D struct {
	Vertex2D
	Z int
}

// 参照を返さないと、メモリ使用量が倍になる？（半分は使用しないという状態に、、、）
func NewVertex3D(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex2D{x, y}, z}
}

func (v3 Vertex3D)Area3D() int {
	return v3.X * v3.Y * v3.Z
}