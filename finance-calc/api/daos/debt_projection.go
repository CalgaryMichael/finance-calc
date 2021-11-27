package daos

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
	debtModels "financeCalc/pkg/debt/models"
)

func buildBulkInsertDebtProjectionStatement(effectiveDate time.Time, projections []*debtModels.DebtProjection) (string, Params) {
	params := Params{
		"effectiveDate": effectiveDate,
	}
	inserts := make([]string, 0, len(projections))
	for i, projection := range projections {
		params[fmt.Sprintf("debtId%d", i)] = projection.Debt.Id
		params[fmt.Sprintf("total%d", i)] = projection.DebtTotal
		params[fmt.Sprintf("paymentSum%d", i)] = projection.PaymentSum
		params[fmt.Sprintf("unappliedSum%d", i)] = projection.UnappliedSum
		insertStatement := fmt.Sprintf("(:debtId%[1]d, :effectiveDate, :total%[1]d, :paymentSum%[1]d, :unappliedSum%[1]d)", i)
		inserts = append(inserts, insertStatement)
	}
	statement := fmt.Sprintf(
		`INSERT INTO finance.debt_projection (debt_id, effective_date, total, payment_sum, unapplied_sum) VALUES %s`,
		strings.Join(inserts, ", "),
	)
	return statement, params
}

func CreateDebtProjections(tx *sqlx.Tx, effectiveDate time.Time, projections []*debtModels.DebtProjection) {
	statement, params := buildBulkInsertDebtProjectionStatement(effectiveDate, projections)
	_, err := tx.NamedExec(statement, params)
	utils.CheckError(err)
}

func GetDebtProjectionsForScenario(tx *sqlx.Tx, scenarioId int) []*debtModels.DebtProjection {
	statement := `
		SELECT
			d.id,
			d.name,
			d.total AS debt_total,
			projection.effective_date,
			projection.total,
			projection.payment_sum,
			projection.unapplied_sum
		FROM finance.debt_projection projection
		JOIN finance.debt d ON d.id = projection.debt_id
		JOIN finance.scenario s ON s.id = d.scenario_id
		WHERE s.id = :scenarioId
	`
	params := Params{
		"scenarioId": scenarioId,
	}
	rows, err := tx.NamedQuery(statement, params)
	utils.CheckError(err)

	var projections []*debtModels.DebtProjection
	for rows.Next() {
		var debt debtModels.Debt
		var projection debtModels.DebtProjection
		err := rows.Scan(
			&debt.Id,
			&debt.DebtName,
			&debt.DebtTotal,
			&projection.EffectiveDate,
			&projection.DebtTotal,
			&projection.PaymentSum,
			&projection.UnappliedSum,
		)
		utils.CheckError(err)
		projection.Debt = &debt
		projections = append(projections, &projection)
	}
	return projections
}
