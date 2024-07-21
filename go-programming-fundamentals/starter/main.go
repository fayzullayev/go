package main

import "fmt"

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehicle
	BodyStyle string
	Year      int
}

func main() {

	var myCar Car = Car{
		Vehicle: Vehicle{
			Make:  "gerg",
			Model: "gerg",
			Year:  0,
		},
		BodyStyle: "dasfasf",
		Year:      2000,
	}

	fmt.Println(myCar.Vehicle.Year)

}
