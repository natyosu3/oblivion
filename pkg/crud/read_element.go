package crud

import (
	"database/sql"
	"oblivion/pkg/error_handler"
	"oblivion/pkg/user"
)

func GetListElement(userid string) ([]user.Element, error) {
	db := Connect()
	defer db.Close()

	var (
		sqlStatement string
		rows         *sql.Rows
		err          error
	)

	sqlStatement = `SELECT id, name, content, remind FROM "Element" WHERE userid = $1`
	rows, err = db.Query(sqlStatement, userid)
	if err != nil {
		return nil, error_handler.SelectError{Message: "Select Error"}
	}

	var elements []user.Element
	for rows.Next() {
		var element user.Element
		if err := rows.Scan(&element.Id, &element.Name, &element.Content, &element.Remind); err != nil {
			return nil, error_handler.SelectError{Message: "Select Error"}
		}
		elements = append(elements, element)
	}

	return elements, nil
}