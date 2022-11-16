package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["yellow"] = "laksjdf"
	printMap(colors)
	println(" ")
	delete(colors, "yellow")

	printMap(colors)

	var colors1 map[string]string
	// colors1["white"] = "#ffffff" -> Não funciona

	printMap(colors1)
	colors2 := make(map[int]string)
	colors2[1] = "Azul"
	for key, value := range colors2 {
		println(key, value)
	}

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
