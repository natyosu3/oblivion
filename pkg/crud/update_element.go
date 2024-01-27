package crud

import (
	"oblivion/pkg/error_handler"
)

func UpdateElement(id string) error {
	db := Connect()
	defer db.Close()

	var (
		sqlStatement string
		err          error
	)

	sqlStatement = `UPDATE "Element" SET remind = $1 WHERE id = $2`
	_, err = db.Exec(sqlStatement,  id)
	if err != nil {
		return error_handler.UpdateError{Message: "Update Error"}
	}

	return nil
}