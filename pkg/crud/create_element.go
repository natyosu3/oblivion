package crud

import (
	"oblivion/pkg/error_handler"
	"oblivion/pkg/utils/general"
)

func InsertElement(userid string, name string, content string, remind string) error {
	db := Connect()
	defer db.Close()

	elementid := general.MakeRandomId()

	sqlStatement := `INSERT INTO "Element" (id, userid, name, content, remind, frequency) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, elementid, userid, name, content, remind, 0)
	if err != nil {
		return error_handler.InsertError{Message: "Insert Error"}
	}

	return nil
}