package main

import "fmt"

//passing array to other functioons
func main() {
	arr := [3]int{1, 2, 3}
	var a int
	a = other(arr, 6)
	fmt.Println(a)
}
func other(b [3]int, size int) int {
	var c = 9
	add := size + c
	return add
}
