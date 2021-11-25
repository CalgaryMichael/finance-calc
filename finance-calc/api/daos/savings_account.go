package daos

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
	savingsModels "financeCalc/pkg/savings/models"
)

func CreateSavingsAccount(tx *sqlx.Tx, scenarioId int, savingsAccount *savingsModels.SavingsAccount) int {
	statement := `
		INSERT INTO finance.savings_account (scenario_id, name, apy, initial_capital, projected_date)
		VALUES (:scenarioId, :name, :apy, :initialCapital, :projectedDate)
		ON CONFLICT DO NOTHING
		RETURNING id
	`
	params := Params{
		"scenarioId":     scenarioId,
		"name":           savingsAccount.Name,
		"apy":            savingsAccount.APY,
		"initialCapital": savingsAccount.InitialCapital,
		"projectedDate":  savingsAccount.ProjectedDate,
	}
	rows, err := tx.NamedQuery(statement, params)
	utils.CheckError(err)

	id, ok := GetInsertedId(rows)
	if !ok {
		panic(errors.New(fmt.Sprintf("Unable to insert SavingsAccount \"%s\"", savingsAccount.Name)))
	}
	return id
}
