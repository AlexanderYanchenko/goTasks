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

func (k Truck) transportationDevice() string {
	return fmt.Sprintln("Truck have a ", k.colors, "and", k.doors, "doors")
}
func (j Sedan) transportationDevice() string {
	return fmt.Sprintln("Sedan have a ", j.colors, "and", j.doors, "doors")
}

func main() {
	t1 := Truck{
		Vehicle{
			2,
			"grey",
		},
		true,
	}
	s1 := Sedan{
		Vehicle{
			4,
			"blue",
		},
		true,
	}

	fmt.Println(t1)
	fmt.Println(t1.colors)
	fmt.Println(t1.transportationDevice())
	fmt.Println(s1)
	fmt.Println(s1.Vehicle)
	fmt.Println(s1.transportationDevice())

}
