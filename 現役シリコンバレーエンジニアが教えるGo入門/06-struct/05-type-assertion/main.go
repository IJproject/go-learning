package main

import "fmt"

func do (i interface{}) {
	// タイプアサーション
	// i.(int), str.(string) など
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T", v)
	}		
}

func main() {
	do(10)      // => 20
	do("Bob")   // => Bob!
	do(true)    // => I don't know bool
}