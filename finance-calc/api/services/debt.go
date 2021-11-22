package services

import (
	"financeCalc/api/daos"
	debtModels "financeCalc/pkg/debt/models"

	"github.com/jmoiron/sqlx"
)

func CreateDebts(tx *sqlx.Tx, scenarioId int, debts []*debtModels.Debt) {
	// TODO: figure out how to do this in two SQL calls
	for _, debt := range debts {
		debtId := daos.CreateDebt(tx, scenarioId, debt)
		CreateDebtPayments(tx, debtId, debt.Payments)
	}
}
