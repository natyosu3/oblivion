package crud

import (
	"fmt"
	"oblivion/pkg/error_handler"
)

func DeleteElement(userid string, elementId string) (error) {
	db := Connect()
	defer db.Close()

	fmt.Println(userid, elementId)

	sqlStatement := `DELETE FROM "Element" WHERE userid = $1 AND id = $2`
	_, err := db.Exec(sqlStatement, userid, elementId)
	if err != nil {
		return error_handler.DeleteError{Message: "Delete Error"}
	}
	return nil
}