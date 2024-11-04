package main

import "fmt"

func main() {
	// 1 Way : Create a map with string keys and string values
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// colors["black"] = "#000000"
	// delete(colors, "white")

	// 2 Way : Create a map with string keys and string values
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#0000ff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
