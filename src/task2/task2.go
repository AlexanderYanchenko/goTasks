package main

import "fmt"

func main() {
	simpleMap := map[string]int{"Alex": 21, "Oleh": 21}
	for key, _ := range simpleMap {
		fmt.Println(key)
	}
	for key, value := range simpleMap {
		fmt.Println(key, value)
	}
}
