//https://golangbot.com/if-statement/

package main

import "fmt"

func main() {
	countries := []string{"USA", "Singapore", "UK", "India", "France", "Indonesia"}
	neededCountries := countries[:len(countries)-2]
	fmt.Println(neededCountries, len(neededCountries), cap(neededCountries))
	copyiedCountries := make([]string, len(neededCountries))
	fmt.Println(copyiedCountries, len(copyiedCountries), cap(copyiedCountries))
	copy(copyiedCountries, neededCountries)
	fmt.Println(copyiedCountries, len(copyiedCountries), cap(copyiedCountries))
}
