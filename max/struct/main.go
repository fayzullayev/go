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

func (u *user) outPutUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() {
	u.firstName = "xxx"
	u.lastName = "xxx"
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

	fmt.Printf("%+v\n", appUser)

	appUser.clearUserName()

	fmt.Printf("%+v", appUser)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
