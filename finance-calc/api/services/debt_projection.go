package services

import (
	"time"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/daos"
	debtModels "financeCalc/pkg/debt/models"
)

func CreateDebtProjections(tx *sqlx.Tx, effectiveDate time.Time, projections []*debtModels.DebtProjection) {
	daos.CreateDebtProjections(tx, effectiveDate, projections)
}

func GetDebtProjectionsForScenario(tx *sqlx.Tx, scenarioId int) []*debtModels.DebtProjection {
	return daos.GetDebtProjectionsForScenario(tx, scenarioId)
}
