package crud

import (
	"database/sql"
	"oblivion/pkg/error_handler"
	"oblivion/pkg/model"
)

func GetElement(id string) (model.Element, error) {
	db := Connect()
	defer db.Close()

	var (
		sqlStatement string
		element      model.Element
		err          error
	)

	sqlStatement = `SELECT id, name, content, remind, frequency FROM "Element" WHERE id = $1`
	err = db.QueryRow(sqlStatement, id).Scan(&element.Id, &element.Name, &element.Content, &element.Remind, &element.Frequency)
	if err != nil {
		return model.Element{}, error_handler.SelectError{Message: err.Error()}
	}

	return element, nil
}

func GetListElement(userid string) ([]model.Element, error) {
	db := Connect()
	defer db.Close()

	var (
		sqlStatement string
		rows         *sql.Rows
		err          error
	)

	sqlStatement = `SELECT id, name, content, remind, frequency FROM "Element" WHERE userid = $1`
	rows, err = db.Query(sqlStatement, userid)
	if err != nil {
		return nil, error_handler.SelectError{Message: "Select Error"}
	}

	var elements []model.Element
	for rows.Next() {
		var element model.Element
		if err := rows.Scan(&element.Id, &element.Name, &element.Content, &element.Remind, &element.Frequency); err != nil {
			return nil, error_handler.SelectError{Message: "Select Error"}
		}
		elements = append(elements, element)
	}

	return elements, nil
}
