package user

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

func (u *User) OutPutUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.age)
}

func (u *User) ClearUserName() {
	u.firstName = "xxx"
	u.lastName = "xxx"
}
