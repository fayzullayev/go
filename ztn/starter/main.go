package main

import "fmt"

func main() {
	routes := []string{"Grocery", "Department Store", "Salon"}

	printSlice("Route 1", routes)

	routes = append(routes, "Home")
	printSlice("Route 2", routes)
	fmt.Println()

	fmt.Println(routes[0], "visited")
	fmt.Println(routes[1], "visited")

	routes = routes[2:]
	printSlice("Remaining locations", routes)

}

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
