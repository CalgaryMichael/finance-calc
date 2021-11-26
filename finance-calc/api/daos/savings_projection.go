package daos

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
	savingsModels "financeCalc/pkg/savings/models"
)

func buildBulkInsertSavingsProjectionStatement(effectiveDate time.Time, projections []*savingsModels.SavingsProjection) (string, Params) {
	params := Params{
		"effectiveDate": effectiveDate,
	}
	inserts := make([]string, 0, len(projections))
	for i, projection := range projections {
		params[fmt.Sprintf("savingsAccountId%d", i)] = projection.SavingsAccount.Id
		params[fmt.Sprintf("total%d", i)] = projection.SavingsTotal
		params[fmt.Sprintf("paymentSum%d", i)] = projection.PaymentSum
		insertStatement := fmt.Sprintf("(:savingsAccountId%[1]d, :effectiveDate, :total%[1]d, :paymentSum%[1]d)", i)
		inserts = append(inserts, insertStatement)
	}
	statement := fmt.Sprintf(
		`INSERT INTO finance.savings_projection (savings_account_id, effective_date, total, payment_sum) VALUES %s`,
		strings.Join(inserts, ", "),
	)
	return statement, params
}

func CreateSavingsProjections(tx *sqlx.Tx, effectiveDate time.Time, projections []*savingsModels.SavingsProjection) {
	statement, params := buildBulkInsertSavingsProjectionStatement(effectiveDate, projections)
	_, err := tx.NamedExec(statement, params)
	utils.CheckError(err)
}
