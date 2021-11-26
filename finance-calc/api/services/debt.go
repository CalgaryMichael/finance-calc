package services

import (
	"github.com/jmoiron/sqlx"

	"financeCalc/api/daos"
	debtModels "financeCalc/pkg/debt/models"
)

func CreateDebts(tx *sqlx.Tx, scenarioId int, debts []*debtModels.Debt) {
	// TODO: figure out how to do this in two SQL calls
	for _, debt := range debts {
		CreateDebt(tx, scenarioId, debt)
	}
}

func CreateDebt(tx *sqlx.Tx, scenarioId int, debt *debtModels.Debt) {
	debt.Id = daos.CreateDebt(tx, scenarioId, debt)
	CreateDebtPayments(tx, debt.Id, debt.Payments)
}
