package main

import "fmt"

func main() {
	// Type 1
	// var colors map[string]string
	
	// Type 2
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)

	// Type 3
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white") // delete a key
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}