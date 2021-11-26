package orchestrators

import (
	"log"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/db"
	"financeCalc/api/models"
	"financeCalc/api/services"
	"financeCalc/pkg/scenario"
	scenarioModels "financeCalc/pkg/scenario/models"
)

func CreateProjections(scenarioRequest models.ScenarioRequest) []*scenarioModels.Projection {
	log.Println("Saving projection info to DB...")
	projections := scenario.BuildProjections(
		scenarioRequest.Scenario,
	)

	db.WithTransaction(func(tx *sqlx.Tx) {
		services.CreateProjections(tx, projections)
	})

	return projections
}
