package main

import "fmt"

type Person struct {
	Name string
	Age int
}

// Stringer
// StringメソッドはStringerとして使用し、出力方法を変更することができる
func (p Person) String() string {
	return fmt.Sprintf("My name is %v.", p.Name)
}

func main() {
	Bob := Person{"Bob", 22}
	fmt.Println(Bob)    // => My name is Bob.
}