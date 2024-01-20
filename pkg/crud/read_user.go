package crud

import (
)

func GetUserId(username string, passhash string) (string, error) {
	db := Connect()
	defer db.Close()

	// 入力されたIDとパスワードが一致した場合, ユーザIDを返す
	sqlStatement := `SELECT id FROM "User" WHERE username = $1 AND password = $2`
	row := db.QueryRow(sqlStatement, username, passhash)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}