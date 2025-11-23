package main

import (
	"fmt"
)

func init() {
	fmt.Println("This is practice.go in statement directory.")
}

func main() {
	Q1()

	result := Q2()
	fmt.Println(result)
}

func Q1() {
	result := 1000
	f := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	for _, v := range f {
		if result > v {
			result = v
		}
	}
	fmt.Println(result) // => 2
}

func Q2() (result int) {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	for _, v := range m {
		result += v
	}
	return
}
