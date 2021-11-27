package services

import (
	"time"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/daos"
	scenarioModels "financeCalc/pkg/scenario/models"
)

func CreateProjections(tx *sqlx.Tx, projections []*scenarioModels.Projection) {
	for _, projection := range projections {
		CreateDebtProjections(tx, projection.EffectiveDate, projection.DebtProjections)
		CreateSavingsProjections(tx, projection.EffectiveDate, projection.SavingsProjections)
	}
}

func GetProjectionDateRangeForScenario(tx *sqlx.Tx, scenarioId int) []time.Time {
	return daos.GetProjectionDateRangeForScenario(tx, scenarioId)
}
