package daos

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
	savingsModels "financeCalc/pkg/savings/models"
)

func buildBulkInsertSavingsPaymentStatement(savingsAccountId int, savingsPayments []*savingsModels.SavingsPayment) (string, Params) {
	params := Params{
		"savingsAccountId": savingsAccountId,
	}
	inserts := make([]string, 0, len(savingsPayments))
	for i, payment := range savingsPayments {
		params[fmt.Sprintf("amount%d", i)] = payment.Amount
		params[fmt.Sprintf("startDate%d", i)] = payment.StartDate
		params[fmt.Sprintf("endDate%d", i)] = payment.EndDate
		insertStatement := fmt.Sprintf("(:savingsAccountId, :amount%[1]d, :startDate%[1]d, :endDate%[1]d)", i)
		inserts = append(inserts, insertStatement)
	}
	statement := fmt.Sprintf(
		`INSERT INTO finance.savings_payment (savings_account_id, amount, start_date, end_date) VALUES %s`,
		strings.Join(inserts, ", "),
	)
	return statement, params
}

func CreateSavingsPayments(tx *sqlx.Tx, savingsAccountId int, savingsPayments []*savingsModels.SavingsPayment) {
	statement, params := buildBulkInsertSavingsPaymentStatement(savingsAccountId, savingsPayments)
	_, err := tx.NamedExec(statement, params)
	utils.CheckError(err)
}
