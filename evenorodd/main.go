package main

import "fmt"

func main() {
	myint := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range myint {
		if v%2 == 0 {
			fmt.Printf("%d is even!\n", v)
		} else {
			fmt.Printf("%d is odd!\n", v)
		}
	}
}
