package model

type Session_model struct {
	SessionId string
	CookieKey string
	User      User
	Token     string
}
