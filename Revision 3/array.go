package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 0
	a[1] = 1 //one way of declaration
	fmt.Println(a[0])
	b := [4]int{1, 2, 3, 4} //short hand
	fmt.Println(b)
	c := [3][3]int{
		{1, 2, 3},
		{4, 5, 6}, //multidimensional array
		{7, 8, 9},
	}
	fmt.Println(c[2][2])
	var arr [1]int
	fmt.Println(arr)    //the default value is 0
	fmt.Println(len(a)) //finding the lenght of array
	happen := [...]string{"indian", "country", "master"}
	fmt.Println(len(happen)) //length of ellipsis is based on the method
	for x := 0; x < len(happen); x++ {
		fmt.Println(happen[x]) //iterating a array
	}
	movies := [...]string{"sazam", "justice league", "sucide squad"}
	Dc := movies
	Dc[2] = "batman dawn of justice"
	fmt.Println(movies)           //this proves array is of value type and not
	fmt.Println(Dc)               //the reference type
	fmt.Println(happen == movies) // comparing two arrays using == operator
	movie := happen               //copying one array to another array
	fmt.Println(movie == happen)
	happens := &happen
	happens[2] = "BLASTER" //this reflects in original array also
	fmt.Println(happen)
	fmt.Println(happens)  
}
