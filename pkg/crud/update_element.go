package crud

import (
	"oblivion/pkg/error_handler"
	"oblivion/pkg/user"
)

func UpdateElement(element user.Element) error {
	db := Connect()
	defer db.Close()

	var (
		sqlStatement string
		err          error
	)

	sqlStatement = `UPDATE "Element" SET name = $1, content = $2, remind = $3, frequency = $4 WHERE id = $5`
	_, err = db.Exec(sqlStatement, element.Name, element.Content, element.Remind, element.Frequency, element.Id)
	if err != nil {
		return error_handler.UpdateError{Message: err.Error()}
	}

	return nil
}
