package services

import (
	"financeCalc/api/daos"
	savingsModels "financeCalc/pkg/savings/models"

	"github.com/jmoiron/sqlx"
)

func CreateSavingsPayments(tx *sqlx.Tx, savingsAccountId int, payments []*savingsModels.SavingsPayment) {
	daos.CreateSavingsPayments(tx, savingsAccountId, payments)
}
