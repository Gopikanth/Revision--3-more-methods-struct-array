package main

import "fmt"

func main() {
	Tv := struct {
		model string
		year  int
	}{
		model: "Bravia",  
		year:  2018,           //an Anonymous structure
	}
	fmt.Println(Tv)
}
