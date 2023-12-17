package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName, lastName, birthDate string
	age                            int
	createdAt                      time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}

	fmt.Println()

	// ... do something awesome with that gathered data!

	outPutUserDetails(&appUser)
}

func outPutUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
