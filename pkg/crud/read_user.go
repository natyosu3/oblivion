package crud

import (
)

func GetUserId(username string) (string, error) {
	db := Connect()
	defer db.Close()

	// 入力されたIDとパスワードが一致した場合, ユーザIDを返す
	sqlStatement := `SELECT userid FROM "User" WHERE username = $1`
	row := db.QueryRow(sqlStatement, username)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func GetPasswordHash(username string) (string, error) {
	db := Connect()
	defer db.Close()

	// ユーザ名からパスワードハッシュを取得
	sqlStatement := `SELECT password FROM "User" WHERE username = $1`
	row := db.QueryRow(sqlStatement, username)

	var passhash string
	err := row.Scan(&passhash)
	if err != nil {
		return "", err
	}

	return passhash, nil
}