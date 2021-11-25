package services

import (
	"financeCalc/api/daos"
	savingsModels "financeCalc/pkg/savings/models"

	"github.com/jmoiron/sqlx"
)

func CreateSavingsAccounts(tx *sqlx.Tx, scenarioId int, savingsAccounts []*savingsModels.SavingsAccount) {
	// TODO: figure out how to do this in two SQL calls
	for _, savingsAccount := range savingsAccounts {
		savingsAccountId := daos.CreateSavingsAccount(tx, scenarioId, savingsAccount)
		CreateSavingsPayments(tx, savingsAccountId, savingsAccount.Payments)
	}
}
