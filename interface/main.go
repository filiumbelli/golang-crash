package main

import "fmt"

func main() {
	englishBot := englishBot{}
	spanishBot := spanishBot{}

	printGreetings(englishBot)
	printGreetings(spanishBot)

	square := square{sideLength: 12.5}
	triangle := triangle{height: 4, base: 3}

	printArea(square)
	printArea(triangle)
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}
