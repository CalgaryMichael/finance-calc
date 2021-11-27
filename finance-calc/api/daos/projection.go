package daos

import (
	"time"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
)

func GetProjectionDateRangeForScenario(tx *sqlx.Tx, scenarioId int) []time.Time {
	statement := `
		WITH debt_dates AS (
			SELECT
				DISTINCT(projection.effective_date) as dates
			FROM finance.debt_projection as projection
			JOIN finance.debt d ON d.id = projection.debt_id
			JOIN finance.scenario s ON s.id = d.scenario_id
			WHERE s.id = :scenarioId
		),
		savings_dates AS (
			SELECT
				DISTINCT(projection.effective_date) AS dates
			FROM finance.savings_projection as projection
			JOIN finance.savings_account s ON s.id = projection.savings_account_id
			JOIN finance.scenario sc ON sc.id = s.scenario_id
			WHERE s.id = :scenarioId
		),
		combined AS (
			SELECT dates FROM debt_dates
			UNION
			SELECT dates FROM savings_dates
		)
		SELECT DISTINCT(dates)
		FROM combined;
	`
	params := Params{
		"scenarioId": scenarioId,
	}
	rows, err := tx.NamedQuery(statement, params)
	utils.CheckError(err)

	var dates []time.Time
	for rows.Next() {
		var date time.Time
		err := rows.Scan(&date)
		utils.CheckError(err)
		dates = append(dates, date)
	}
	return dates
}
