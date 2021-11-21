package dao

import (
	"database/sql"

	"financeCalc/api/utils"
	scenarioModels "financeCalc/pkg/scenario/models"
)

func CreateScenario(conn *sql.DB, userId int, scenario scenarioModels.Scenario) {
	statement := `
		INSERT INTO finance.scenario(name, start_date, sort_key, reverse_sort)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	params := []interface{}{
		"scenario",
		scenario.StartDate,
		scenario.SortKey,
		scenario.ReverseSort,
	}
	_, err := conn.Exec(statement, params...)
	utils.CheckError(err)
}
