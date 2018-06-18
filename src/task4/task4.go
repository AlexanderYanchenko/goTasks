package main

import "fmt"

func main() {
	p1 := Person{
		"Yan",
		"Alexandrov",
	}
	fmt.Println(p1)
	fmt.Println(p1.fName)
}

type Person struct {
	fName string
	lName string
}
