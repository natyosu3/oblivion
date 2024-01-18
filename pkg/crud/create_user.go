package crud

import (
	"oblivion/pkg/error_hanndler"
	_ "github.com/lib/pq"
)


func checkExsistUser(name string, email string) (bool, error) {
	db := Connect()
	defer db.Close()

	sqlStatement := `SELECT username, email FROM "User" WHERE username = $1 OR email = $2`
	row := db.QueryRow(sqlStatement, name, email)

	var username string
	var mail string
	err := row.Scan(&username, &mail)
	if err != nil {
		return false, nil
	}

	return true, error_hanndler.AlreadyExsistUserError{Message: "User already exsist"}
}


func InsertUser(name string, email string, password string) error {
	db := Connect()
	defer db.Close()

	// 既に登録されているユーザー名かメールアドレスかを確認
	exsist, err := checkExsistUser(name, email)
	if exsist {
		return err
	}


	sqlStatement := `INSERT INTO "User" (username, email, password) VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, name, email, password)

	if err != nil {
		return err
	}

	return nil
}