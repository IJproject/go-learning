package main

import "fmt"

func main() {
	var n int = 100
	fmt.Println(n)   // => 100
	fmt.Println(&n)  // => 0x140000a4008

	var p *int = &n
	fmt.Println(p)   // => 0x140000a4008
	fmt.Println(*p)  // => 100

	// 値渡し
	oneValue(n)
	fmt.Println(n)   // => 100

	// 参照渡し
	onePointer(&n)
	fmt.Print(n)     // => 1

	// メモリ領域のみ確保する（newはポインタを返す）
	var p1 *int = new(int)
	var p2 *int
	fmt.Println(p1)  // => 10x1400000e0d0
	fmt.Println(p2)  // => <nil>
}

func oneValue(x int) {
	x = 1
}

func onePointer(x *int) {
	*x = 1
}