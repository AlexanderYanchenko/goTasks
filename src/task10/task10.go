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

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Sprintln(t.transportationDevice)
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

	report(t1)
	report(s1)

}
