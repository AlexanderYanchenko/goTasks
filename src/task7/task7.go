package main

import "fmt"

type Vehicle struct {
	doors  int
	colors string
}

type Truck struct {
	Vehicle
	fourWhell bool
}

type Sedan struct {
	Vehicle
	luxury bool
}

func main() {
	t1 := Truck{
		Vehicle{
			2,
			"blue",
		},
		true,
	}
	s1 := Sedan{
		Vehicle{
			4,
			"red",
		},
		true,
	}

	fmt.Println(t1)
	fmt.Println(t1.colors)
	fmt.Println(s1)
	fmt.Println(s1.Vehicle)

}
