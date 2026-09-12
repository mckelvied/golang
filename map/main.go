package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"
	delete(colors, "red")

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}
