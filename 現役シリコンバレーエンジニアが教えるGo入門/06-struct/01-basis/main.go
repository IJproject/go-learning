package main

import "fmt"

type Vertex struct {
	X int
	Y int
	// X, Y int
	// 小文字はプライベート（同パッケージ内からはアクセス可能）
}

func main() {
	v1 := Vertex{ X: 1, Y: 2 }
	fmt.Println(v1)          // => {1 2}
	fmt.Println(v1.X, v1.Y)  // => 1 2

	v1.X = 100
	fmt.Println(v1.X, v1.Y)  // => 100 2

	v2 := Vertex{ X: 1 }     // Yにはintの初期値が入る
	fmt.Println(v2)          // => {1 0} 

	var v3 Vertex
	fmt.Println(v3)          // => {0 0}

	v4 := new(Vertex)
	fmt.Println(v4)          // => &{0 0}

	v5 := &Vertex{}
	fmt.Println(v5)          // => &{0 0}

	v11 := Vertex{1, 2}
	changeVertexValue(v11)
	fmt.Println(v11)         // => {1 2}

	v12 := &Vertex{1, 2}
	changeVertexPointer(v12)
	fmt.Println(*v12)         // => {1000 2}
}

func changeVertexValue(v Vertex){
	v.X = 1000
}

func changeVertexPointer(v *Vertex){
	v.X = 1000
	// (*v).X = 1000
}