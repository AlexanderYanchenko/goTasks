package main

import "fmt"

func main() {
	p1 := Person{
		"Yan",
		"Alexandrov",
	}
	s := p1.walk()
	fmt.Println(s)
}

type Person struct {
	fName string
	lName string
}

func (p Person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}
