package user

import (
)

type User struct {
	UserId   	 string
	UserName 	 string
	EmailAddress string
	Password 	 string
	Comportement Comportement
}

type Comportement struct {
	Id string
	Elements []Element
}

type Element struct {
	Id 	 int
	Name string
	Content string
	Priod_info Priod_info
}

type Priod_info struct {
	RegistrationDate string
	NextRemindDate   string
}