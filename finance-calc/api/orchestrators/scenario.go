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

func CreateScenario(scenarioRequest models.ScenarioRequest) []*scenarioModels.Projection {
	log.Println("Saving scenario info to DB...")
	projections := scenario.BuildProjections(
		scenarioRequest.Scenario,
	)

	db.WithTransaction(func(tx *sqlx.Tx) {
		// TODO: actually capture the user id for the request
		scenarioId := services.CreateScenario(tx, 0, scenarioRequest.Scenario)
		services.CreateDebts(tx, scenarioId, scenarioRequest.Scenario.Debts)
		services.CreateSavingsAccounts(tx, scenarioId, scenarioRequest.Scenario.SavingsAccounts)
	})

	// TODO: save projections to DB

	return projections
}
