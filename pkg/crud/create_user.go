package crud

import (
	"database/sql"
	"oblivion/pkg/error_handler"
	"oblivion/pkg/utils/general"

	_ "github.com/lib/pq"
)


func checkExsistUser(name string, email string) (bool, error) {
	db := Connect()
	defer db.Close()

	var (
		sqlStatement string
		row *sql.Row
		username string
		mail string
		err error
	)

	if email != "" {
		sqlStatement = `SELECT username, email FROM "User" WHERE username = $1 OR email = $2`
		row = db.QueryRow(sqlStatement, name, email)
		err = row.Scan(&username, &mail)
		if err != nil {
			return false, nil
		}
	} else {
		sqlStatement = `SELECT username FROM "User" WHERE username = $1`
		row = db.QueryRow(sqlStatement, name)
		err = row.Scan(&username)
		if err != nil {
			return false, nil
		}
	}

	return true, error_handler.AlreadyExsistUserError{Message: "User already exsist"}
}


func InsertUser(name string, email string, password string) error {
	db := Connect()
	defer db.Close()

	// 既に登録されているユーザー名かメールアドレスかを確認
	exsist, err := checkExsistUser(name, email)
	if exsist {
		return err
	}

	userid := general.MakeRandomId()

	// emailが空の場合はNULLとして扱う
	var emailValue interface{}
	if email == "" {
		emailValue = nil
	} else {
		emailValue = email
	}

	sqlStatement := `INSERT INTO "User" (userid, username, email, password) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, userid, name, emailValue, password)

	if err != nil {
		return err
	}

	return nil
}