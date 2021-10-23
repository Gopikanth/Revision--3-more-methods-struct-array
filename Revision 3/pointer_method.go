package main

import "fmt"

type detail struct {
	name string
	age  int
}

func main() {
	b := detail{name: "kana",
		age: 25,
	}
	fmt.Println(b.name)
	fmt.Println(b.age)
	//now after creating pointer
	p :=&b
	p.show("Game")
	fmt.Println(p.name)
}
func (a *detail) show( other  string) {
	(*a).name = other
 }
