package dao

import (
	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
	scenarioModels "financeCalc/pkg/scenario/models"
)

func CreateScenario(conn *sqlx.DB, userId int, scenario scenarioModels.Scenario) {
	statement := `
		INSERT INTO finance.scenario(name, start_date, sort_key, reverse_sort)
		VALUES (:scenario, :startDate, :sortKey, :reverseSort)
		RETURNING id;
	`
	params := map[string]interface{}{
		"scenario":    "scenario",
		"startDate":   scenario.StartDate,
		"sortKey":     scenario.SortKey,
		"reverseSort": scenario.ReverseSort,
	}
	_, err := conn.NamedExec(statement, params)
	utils.CheckError(err)
}
