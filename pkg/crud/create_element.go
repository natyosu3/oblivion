package crud

import (
	"oblivion/pkg/error_handler"
	"oblivion/pkg/utils/general"
)

func InsertElement(userid string, name string, content string) error {
	db := Connect()
	defer db.Close()

	elementid := general.MakeRandomId()

	sqlStatement := `INSERT INTO "Element" (id, userid, name, content) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, elementid, userid, name, content)
	if err != nil {
		return error_handler.InsertError{Message: "Insert Error"}
	}

	return nil
}