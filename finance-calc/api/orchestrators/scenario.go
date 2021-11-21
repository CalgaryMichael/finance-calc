package orchestrators

import (
	"log"

	daos "financeCalc/api/daos"
	"financeCalc/api/db"
	"financeCalc/api/models"
	"financeCalc/pkg/scenario"

	scenarioModels "financeCalc/pkg/scenario/models"
)

func CreateScenario(scenarioRequest models.ScenarioRequest) []*scenarioModels.Projection {
	log.Println("Saving scenario info to DB...")
	projections := scenario.BuildProjections(
		scenarioRequest.Scenario,
	)

	conn := db.GetDbInstance()
	// TODO: actually capture the user id for the request
	daos.CreateScenario(conn, 0, scenarioRequest.Scenario)

	return projections
}
