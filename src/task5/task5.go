package main

import "fmt"

func main() {
	p1 := Person{
		"Yan",
		"Alexandrov",
		[]string{"Мандарин", "Апельсин", "груша"},
	}
	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)
	for key, value := range p1.favFood {
		fmt.Println(key, value)
	}
}

type Person struct {
	fName   string
	lName   string
	favFood []string
}
