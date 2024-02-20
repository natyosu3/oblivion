module github.com/boj/redistore

go 1.21.4

require (
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.1.1
)

require github.com/gorilla/context v1.1.1 // indirect

replace github.com/gomodule/redigo => ../redigo
