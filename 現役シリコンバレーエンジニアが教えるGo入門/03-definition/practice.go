package main

import (
	"fmt"
)

func init() {
	fmt.Println("This is practice.go in definition directory.")
}

func main() {
	Q1()
	Q2()
	Q3()
}

func Q1() {
	f := 1.11
	f_int := int(f)
	fmt.Println(f_int)  // => 1
}

func Q2() {
	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])  // => [5 6]
}

func Q3() {
	m := map[string]int{
		"Mike": 20,
		"Nancy": 24,
		"Messi": 30
	}
	fmt.Printf("%T %v\n", m, m)
	mike, ok := m["Mike"]
	fmt.Println(mike, ok)  // => 20 true
}
