package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "#ff0000",
		"green":  "#00ff00",
		"blue":   "#0000ff",
		"yellow": "#ffff00",
		"purple": "#00ff00",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// Won't delete this first main() for notes posterity
//func main() {
//	// Syntax 1 - with var keyword, if we need to figure out what values to add later
//	var colors map[string]string
//	colors = make(map[string]string)
//
//	// Syntax 2 - with 'make' built-in fxn (exact equivalent to syntax 1)
//	colorsAlt := make(map[string]string)
//
//	// Syntax 2 - but with inlined value initialization
//	colorsAlt2 := map[string]string{
//		"red":    "#ff0000",
//		"green":  "#00ff00",
//		"blue":   "#0000ff",
//		"yellow": "#ffff00",
//		"purple": "#00ff00",
//	}
//
//	colors["black"] = "000000"
//	colorsAlt["white"] = "#ffffff"
//
//	// Remove entries from a map
//	delete(colors, "black")
//	delete(colorsAlt2, "yellow")
//
//	fmt.Println(colors)
//	fmt.Println(colorsAlt)
//	fmt.Println(colorsAlt2)
//}
