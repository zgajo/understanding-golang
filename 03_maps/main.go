package main

import (
	"fmt"
)

func main() {
	var colors = map[string]string{
		"red":   "#ff0000",
		"white": "#fff",
	}

	colors["green"] = "#32sdds"

	delete(colors, "white")

	fmt.Println(colors)
}
