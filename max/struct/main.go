package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var err error

	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')

	if err != nil {
		return
	}

	fmt.Println("readString", readString)

	//firstName := getUserData("Please enter your first name: ")
	//lastName := getUserData("Please enter your last name: ")
	//birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//var appUser *user.User

	//appUser := user.NewAdmin(firstName, lastName)

	//fmt.Printf("%+v", appUser.)

	//appUser = &user.User{
	//	FirstName: firstName,
	//	LastName:  lastName,
	//	BirthDate: birthDate,
	//}

	//if err != nil {
	//	panic("Error -- " + err.Error())
	//}
	//
	//appUser.OutPutUserDetails()
	//
	//appUser.ClearUserName()
	//
	//appUser.OutPutUserDetails()
}
