package main

import "fmt"

type mobile struct {
	string
	int
	float32
}

func main() {
	description := mobile{
		"poco",
		23000,
		6.5,
	}
	fmt.Println(description.string) //accessing a particular type
	fmt.Println(description) //displaying the structure
}