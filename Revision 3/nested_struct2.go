package main

import "fmt"

type teacher struct {
	name    string
	subject string
	detail  student
}
type student struct {
	name    string
	subject string
}

func main() {
	all := teacher{
		name:    "raju",
		subject: "maths",
		detail:  student{"jaggu", "english"},
	}
	fmt.Println(all.name)        //displays name of teacher
	fmt.Println(all.detail.name) //displays name of student

}
