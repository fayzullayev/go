package main

import "fmt"

func main() {
	//var a = make(map[string]string)
	//a["brand"] = "Ford"
	//a["model"] = "Mustang"
	//a["year"] = "1964"
	//
	////val, ok := a["years"]
	//
	//fmt.Println(a["year"])
	//
	//var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}
	//
	//for k, v := range a {
	//	fmt.Println(k, v)
	//}

	//a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	//
	//for k, v := range a {
	//	fmt.Printf("%v : %v, ", k, v)
	//}

	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b []string // defining the order

	b = append(b, "one", "two", "three", "four")

	for k, v := range a { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println(b)
	fmt.Println()

	for _, element := range b { // loop with the defined order
		fmt.Printf("%v : %v, ", element, a[element])
	}

}
