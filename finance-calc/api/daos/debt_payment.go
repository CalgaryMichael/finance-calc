package daos

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
	debtModels "financeCalc/pkg/debt/models"
)

func buildBulkInsertDebtPaymentStatement(debtId int, debtPayments []*debtModels.DebtPayment) (string, Params) {
	params := Params{
		"debtId": debtId,
	}
	inserts := make([]string, 0, len(debtPayments))
	for i, payment := range debtPayments {
		params[fmt.Sprintf("amount%d", i)] = payment.Amount
		params[fmt.Sprintf("carryOver%d", i)] = payment.CarryOver
		params[fmt.Sprintf("startDate%d", i)] = payment.StartDate
		params[fmt.Sprintf("endDate%d", i)] = payment.EndDate
		insertStatement := fmt.Sprintf("(:debtId, :amount%[1]d, :carryOver%[1]d, :startDate%[1]d, :endDate%[1]d)", i)
		inserts = append(inserts, insertStatement)
	}
	statement := fmt.Sprintf(
		`INSERT INTO finance.debt_payment (debt_id, amount, carry_over, start_date, end_date) VALUES %s`,
		strings.Join(inserts, ", "),
	)
	return statement, params
}

func CreateDebtPayments(tx *sqlx.Tx, debtId int, debtPayments []*debtModels.DebtPayment) {
	statement, params := buildBulkInsertDebtPaymentStatement(debtId, debtPayments)
	_, err := tx.NamedExec(statement, params)
	utils.CheckError(err)
}
