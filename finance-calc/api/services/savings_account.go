package services

import (
	"github.com/jmoiron/sqlx"

	"financeCalc/api/daos"
	savingsModels "financeCalc/pkg/savings/models"
)

func CreateSavingsAccounts(tx *sqlx.Tx, scenarioId int, savingsAccounts []*savingsModels.SavingsAccount) {
	// TODO: figure out how to do this in two SQL calls
	for _, savingsAccount := range savingsAccounts {
		CreateSavingsAccount(tx, scenarioId, savingsAccount)
	}
}

func CreateSavingsAccount(tx *sqlx.Tx, scenarioId int, savingsAccount *savingsModels.SavingsAccount) {
	savingsAccount.Id = daos.CreateSavingsAccount(tx, scenarioId, savingsAccount)
	CreateSavingsPayments(tx, savingsAccount.Id, savingsAccount.Payments)
}
