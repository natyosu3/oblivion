package crud

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)


func InsertUser(db *sql.DB, name string, email string, password string) {
	sqlStatement := `INSERT INTO "User" (username, email, password) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, name, email, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("New user added")

}