package main

import "fmt"
// Method with Non-Struct Type Receiver
func main() {
	a := multi(3, 4)
	fmt.Println(a)

}
func multi(p1, p2 int) int {
	return p1 * p2
}
