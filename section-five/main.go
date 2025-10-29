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

	fmt.Println(colors)
	fmt.Println(colors["red"])

}
