package main

import "fmt"

func main() {
	Q1()
	Q2()
}

func Q1() {
	i := 10
	p := &i
	fmt.Println("pointer:", p)
	fmt.Println("value:", *p)
}

func Q2() {
	fmt.Println("内容理解しているかというより、めんどくさいだけやな")
}
