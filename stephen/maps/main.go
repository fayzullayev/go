package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff0444",
	}

	fmt.Println(colors)
	colors["red"] = "qwertyui"
	colors["red1"] = "qwertyui222"

	fmt.Println(colors)

	delete(colors, "red1fwefwefwe")

	fmt.Println(colors)

	dd, ok := colors["red11"]

	fmt.Println(dd, ok)

	fmt.Println("--------------------------------------------")
	for k, v := range colors {
		fmt.Println(k, v)
	}
	fmt.Println("--------------------------------------------")
	change(colors)
	for k, v := range colors {
		fmt.Println(k, v)
	}
	fmt.Println("--------------------------------------------")
}

func change(c map[string]string) {
	c["red1"] = "hello"
}
