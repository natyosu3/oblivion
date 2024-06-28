package model

type User struct {
	UserId       string
	UserName     string
	EmailAddress string
	Password     string
	Comportement Comportement
}

type Comportement struct {
	Id       string
	Elements []Element
}

type Element struct {
	Id        string
	Name      string
	Content   string
	UserId    string
	Remind    string
	Frequency int
}

type Priod_info struct {
	RegistrationDate string
	NextRemindDate   string
}
