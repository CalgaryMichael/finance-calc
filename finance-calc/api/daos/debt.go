package daos

import (
	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
	debtModels "financeCalc/pkg/debt/models"
)

func CreateDebt(tx *sqlx.Tx, scenarioId int, debt *debtModels.Debt) int {
	statement := `
		INSERT INTO finance.debt (scenario_id, name, total, interest_rate)
		VALUES (:scenarioId, :name, :total, :interestRate)
		ON CONFLICT DO NOTHING
		RETURNING id
	`
	params := Params{
		"scenarioId":   scenarioId,
		"name":         debt.DebtName,
		"total":        debt.DebtTotal,
		"interestRate": debt.InterestRate,
	}
	rows, err := tx.NamedQuery(statement, params)
	utils.CheckError(err)

	var lastInsertedId int
	for rows.Next() {
		err := rows.Scan(&lastInsertedId)
		utils.CheckError(err)

		// we are only trying to insert one value
		// but we will break after the first result just in case
		break
	}

	// manually close since we are not necessarily fully iterating over every row
	rows.Close()

	return lastInsertedId
}
