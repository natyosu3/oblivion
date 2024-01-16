package user

import (
	"fmt"
)

type User struct {
	UserId   	 int
	UserName 	 string
	EmailAddress string
	Password 	 string
	Comportement Comportement
}

type Comportement struct {
	Id int
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


func init() {
	fmt.Println("user init")
	user := User{UserId: 1, UserName: "test", EmailAddress: "test@test.test", Password: "test"}
	user.Comportement = Comportement{Id: 1}
	user.Comportement.Elements = append(user.Comportement.Elements, Element{Id: 1, Name: "test", Content: "test"})
	user.Comportement.Elements = append(user.Comportement.Elements, Element{Id: 2, Name: "test2", Content: "test2"})
	user.Comportement.Elements[0].Priod_info = Priod_info{RegistrationDate: "2018-01-01", NextRemindDate: "2018-01-01"}
	user.Comportement.Elements[1].Priod_info = Priod_info{RegistrationDate: "2018-01-01", NextRemindDate: "2018-01-01"}

	fmt.Println(user)
}

func Usera() {
	fmt.Println("user")
}
