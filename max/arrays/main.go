package main

import "fmt"

func main() {

	//names := make([]string, 5)
	var names = make([]string, 3)
	fmt.Println(names, len(names), cap(names))

	names = append(names, "Mirodil", "Zayab")
	//names = append(names, "Muhammad", "Umar" )

	fmt.Println(names, len(names), cap(names))

}
