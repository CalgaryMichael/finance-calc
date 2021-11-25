package daos

import (
	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
)

func GetInsertedId(rows *sqlx.Rows) (int, bool) {
	insertedId := -1 // some number that the db will never return
	for rows.Next() {
		err := rows.Scan(&insertedId)
		utils.CheckError(err)

		// we are only trying to insert one value
		// but we will break after the first result just in case
		break
	}

	// manually close since we are not necessarily fully iterating over every row
	rows.Close()
	return insertedId, insertedId != -1
}
