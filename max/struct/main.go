package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName, lastName, birthDate string
	age                            int
	createdAt                      time.Time
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("validation error")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		age:       0,
		createdAt: time.Now(),
	}, nil
}

func (u *User) outPutUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) clearUserName() {
	u.firstName = "xxx"
	u.lastName = "xxx"
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := NewUser(firstName, lastName, birthDate)

	if err != nil {
		panic("Error -- " + err.Error())
	}

	fmt.Printf("%+v\n", appUser)

	appUser.clearUserName()

	fmt.Printf("%+v", appUser)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
