package crud

import (
	_ "github.com/lib/pq"
	"fmt"
)



func InsertUser(name string, email string, password string) error {
	db := Connect()
	defer db.Close()

	sqlStatement := `INSERT INTO "User" (username, email, password) VALUES ($1, $2, $3)`
	res, err := db.Exec(sqlStatement, name, email, password)

	fmt.Println(res)
	if err != nil {
		return err
	}
	fmt.Println("New user added")
	return nil
}