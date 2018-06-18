package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for key, _ := range arr {
		fmt.Println("\n", key)
	}
	for key, value := range arr {
		fmt.Println("\n", key, value)
	}

}
