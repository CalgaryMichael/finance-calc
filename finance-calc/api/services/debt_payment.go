package services

import (
	"financeCalc/api/daos"
	debtModels "financeCalc/pkg/debt/models"

	"github.com/jmoiron/sqlx"
)

func CreateDebtPayments(tx *sqlx.Tx, debtId int, debtPayments []*debtModels.DebtPayment) {
	daos.CreateDebtPayments(tx, debtId, debtPayments)
}
