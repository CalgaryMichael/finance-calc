package daos

import (
	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
	scenarioModels "financeCalc/pkg/scenario/models"
)

func CreateScenario(tx *sqlx.Tx, userId int, scenario scenarioModels.Scenario) int {
	statement := `
		INSERT INTO finance.scenario(name, start_date, sort_key, reverse_sort)
		VALUES (:scenario, :startDate, :sortKey, :reverseSort)
		RETURNING id
	`
	params := Params{
		"scenario":    "scenario",
		"startDate":   scenario.StartDate,
		"sortKey":     scenario.SortKey,
		"reverseSort": scenario.ReverseSort,
	}
	rows, err := tx.NamedQuery(statement, params)
	utils.CheckError(err)

	var insertedId int
	for rows.Next() {
		err := rows.Scan(&insertedId)
		utils.CheckError(err)

		// we are only trying to insert one value
		// but we will break after the first result just in case
		break
	}

	// manually close since we are not necessarily fully iterating over every row
	rows.Close()

	return insertedId
}
