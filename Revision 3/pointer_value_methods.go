package main

import "fmt"

type facts struct {
	india    string
	pakistan string
}

func main() {
	again := facts{
		india:    "win",
		pakistan: "loss",
	}
	p := &again //calling with pointer
	p.show("looser")
	fmt.Println(p.pakistan)
	(&again).show_1() //calling with value
	fmt.Println(again.india)

}
func (a *facts) show(other string) {
	(*a).pakistan = other
}
func (a facts) show_1() {
	a.pakistan = "loss"
	fmt.Println(a.pakistan)
}
