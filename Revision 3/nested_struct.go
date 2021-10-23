package main

import "fmt"

type File struct {
	name string
	num  int
}
type Search struct {
	detail File
}

func main() {
	Result := Search{
		detail: File{"Crime Diary", 25},
	}
	fmt.Println(Result)
}