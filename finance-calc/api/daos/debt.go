package daos

import (
	"errors"
	"fmt"

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

	id, ok := GetInsertedId(rows)
	if !ok {
		panic(errors.New(fmt.Sprintf("Unable to insert Debt \"%s\"", debt.DebtName)))
	}
	return id
}
