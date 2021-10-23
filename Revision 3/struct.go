package main

import "fmt"

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {
	a := Address{"Tamilnadu", "Chennai", 69000}
	fmt.Println(a)
	b := Address{"kerla", "kochi", 56747} //calling the same struct but different items
	fmt.Println(b)
	fmt.Println(a.city) // accessing elements in the list
	b.city = "Trichy"   // assigning values to struct
	fmt.Println(b)      //displaying the updated value
	p := &Address{"Delhi", "Agra",34567} //pointer in struct
	fmt.Println((*p).Name) 
	fmt.Println(p.Name) //calling can also be done explicitly
	

}
