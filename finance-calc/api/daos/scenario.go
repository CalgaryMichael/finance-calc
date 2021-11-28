package daos

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
	scenarioModels "financeCalc/pkg/scenario/models"
)

func CreateScenario(tx *sqlx.Tx, userId int, scenario scenarioModels.Scenario) int {
	statement := `
		INSERT INTO finance.scenario(user_id, name, start_date, sort_key, reverse_sort)
		VALUES (:userId, :name, :startDate, :sortKey, :reverseSort)
		RETURNING id
	`
	params := Params{
		"userId":      userId,
		"name":        "scenario",
		"startDate":   scenario.StartDate,
		"sortKey":     scenario.SortKey,
		"reverseSort": scenario.ReverseSort,
	}
	rows, err := tx.NamedQuery(statement, params)
	utils.CheckError(err)

	id, ok := GetInsertedId(rows)
	if !ok {
		panic(errors.New(fmt.Sprintf("Unable to insert Scenario \"%d\"", userId)))
	}
	return id
}

func GetScenarios(tx *sqlx.Tx, userId int) []*scenarioModels.Scenario {
	statement := `
		SELECT
			id,
			start_date,
			sort_key,
			reverse_sort
		FROM finance.scenario
		WHERE scenario.user_id = :userId
	`
	params := Params{
		"userId": userId,
	}
	rows, err := tx.NamedQuery(statement, params)
	utils.CheckError(err)

	var scenarios []*scenarioModels.Scenario
	for rows.Next() {
		var scenario scenarioModels.Scenario
		err := rows.Scan(
			&scenario.Id,
			&scenario.StartDate,
			&scenario.SortKey,
			&scenario.ReverseSort,
		)
		utils.CheckError(err)
		scenarios = append(scenarios, &scenario)
	}
	return scenarios
}
