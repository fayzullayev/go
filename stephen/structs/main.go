package main

import "fmt"

type ContactInfo struct {
	Email string
	Zip   int
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Weight    float64
	IsMarried bool
	ContactInfo
}

func (p *Person) Print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) UpdateName(name string) {
	(*p).FirstName = name
}

func main() {

	var person = Person{
		FirstName: "Mirodil",
		LastName:  "Fayzullayev",
		Age:       28,
		Weight:    75,
		IsMarried: true,
		ContactInfo: ContactInfo{
			Email: "odilfayz22@gmail.com",
			Zip:   100001,
		},
	}

	person.Print()
	person.UpdateName("Odil")
	person.Print()

}
