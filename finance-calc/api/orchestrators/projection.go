package orchestrators

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/db"
	"financeCalc/api/models"
	"financeCalc/api/services"
	debtModels "financeCalc/pkg/debt/models"
	savingsModels "financeCalc/pkg/savings/models"
	"financeCalc/pkg/scenario"
	scenarioModels "financeCalc/pkg/scenario/models"
)

func CreateProjections(scenarioRequest models.ScenarioRequest) {
	log.Println("Saving projection info to DB...")
	projections := scenario.BuildProjections(
		scenarioRequest.Scenario,
	)

	db.WithTransaction(func(tx *sqlx.Tx) {
		services.CreateProjections(tx, projections)
	})
}

func buildProjections(
	dates []time.Time,
	debtProjections []*debtModels.DebtProjection,
	savingsProjections []*savingsModels.SavingsProjection,
) []*scenarioModels.Projection {
	var projections []*scenarioModels.Projection
	for _, date := range dates {
		projection := scenarioModels.Projection{
			EffectiveDate:      date,
			DebtProjections:    debtModels.GetDebtProjectionsForDate(debtProjections, date),
			SavingsProjections: savingsModels.GetSavingsProjectionsForDate(savingsProjections, date),
		}
		projections = append(projections, &projection)
	}
	return projections
}

func GetProjectionsForScenario(projectionRequest models.ProjectionRequest) []*scenarioModels.Projection {
	log.Printf("Getting projections for scenario %d\n", projectionRequest.ScenarioId)

	var dates []time.Time
	var debtProjections []*debtModels.DebtProjection
	var savingsProjections []*savingsModels.SavingsProjection
	db.WithTransaction(func(tx *sqlx.Tx) {
		dates = services.GetProjectionDateRangeForScenario(tx, projectionRequest.ScenarioId)
		debtProjections = services.GetDebtProjectionsForScenario(tx, projectionRequest.ScenarioId)
		savingsProjections = services.GetSavingsProjectionsForScenario(tx, projectionRequest.ScenarioId)
	})
	return buildProjections(dates, debtProjections, savingsProjections)
}
