package main

import "fmt"

func main() {
	colors := map[string]string{}

	// colors := make(map[string]string)

	colors["whit"] = "#FFFFFF"
	printMap(colors)
	changeKey(colors, "whit", "0000000")
	printMap(colors)

}
func changeKey(m map[string]string, k string, v string) {
	m[k] = v
}
func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Color", color, "has hex", hex)
	}
}

type Color struct {
	name string
	hex  string
}
