package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	age       int
	createdAt time.Time
	NickName  string
}

type Admin struct {
	Email    string
	Password string
	User
}

func New(firstName, lastName, birthDate string) (*User, error) {

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

func NewAdmin(email, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "----",
			age:       28,
			createdAt: time.Now(),
		},
	}
}

func (u *User) OutPutUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.age)
}

func (u *User) ClearUserName() {
	u.firstName = "xxx"
	u.lastName = "xxx"
}
