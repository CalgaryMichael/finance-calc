package services

import (
	"github.com/jmoiron/sqlx"

	scenarioModels "financeCalc/pkg/scenario/models"
)

func CreateProjections(tx *sqlx.Tx, projections []*scenarioModels.Projection) {
	for _, projection := range projections {
		CreateDebtProjections(tx, projection.EffectiveDate, projection.DebtProjections)
		CreateSavingsProjections(tx, projection.EffectiveDate, projection.SavingsProjections)
	}
}
