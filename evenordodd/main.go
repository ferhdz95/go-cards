package main

import (
	"fmt"
)

func main() {
	slice := []int{}
	for i := 0; i < 11; i++ {
		slice = append(slice, i)
	}
	for _, s := range slice {
		if s%2 != 0 {
			fmt.Println(s, " is odd")
		} else {
			fmt.Println(s, " is even")
		}
	}
}
