package crud

import (
	"oblivion/pkg/error_handler"
	"oblivion/pkg/user"
)

func UpdateElement(id string, element user.Element, nextday string) error {
	db := Connect()
	defer db.Close()

	var (
		sqlStatement string
		err          error
	)

	element.Frequency = element.Frequency + 1

	sqlStatement = `UPDATE "Element" SET remind = $1, frequency = $2 WHERE id = $3`
	_, err = db.Exec(sqlStatement, nextday, element.Frequency, id)
	if err != nil {
		return error_handler.UpdateError{Message: "Update Error"}
	}

	return nil
}