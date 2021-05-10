package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"
	// colors[10] = "#fffff"

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	// delete(colors, 10)
	// fmt.Println(colors)
	// fmt.Println(colors[10])
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
