package main

import (
	"github.com/structs/user"
)

func main() {
	var err error

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthDate)

	//appUser = &user.User{
	//	FirstName: firstName,
	//	LastName:  lastName,
	//	BirthDate: birthDate,
	//}

	if err != nil {
		panic("Error -- " + err.Error())
	}

	appUser.OutPutUserDetails()

	appUser.ClearUserName()

	appUser.OutPutUserDetails()

}
