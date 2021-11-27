package services

import (
	"time"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/daos"
	savingsModels "financeCalc/pkg/savings/models"
)

func CreateSavingsProjections(tx *sqlx.Tx, effectiveDate time.Time, projections []*savingsModels.SavingsProjection) {
	daos.CreateSavingsProjections(tx, effectiveDate, projections)
}

func GetSavingsProjectionsForScenario(tx *sqlx.Tx, scenarioId int) []*savingsModels.SavingsProjection {
	return daos.GetSavingsProjectionsForScenario(tx, scenarioId)
}
