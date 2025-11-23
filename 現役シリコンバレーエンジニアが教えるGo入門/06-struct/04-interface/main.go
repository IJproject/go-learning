// ダックタイピング：オブジェクトが特定の型であるかどうかではなく、そのオブジェクトが持つメソッドや属性（振る舞い）に基づいて型を判断するプログラミングの概念
// もしアヒルのように歩き、アヒルのように鳴くなら、それはアヒルである
package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get Out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	DriveCar(mike)  // Mr.Mike, Run
	DriveCar(x)     // Mr.X, Get Out
}