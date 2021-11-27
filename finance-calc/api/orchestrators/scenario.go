package orchestrators

import (
	"log"
	"sync"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/db"
	"financeCalc/api/models"
	"financeCalc/api/services"
)

func CreateScenario(scenarioRequest models.ScenarioRequest) int {
	log.Println("Saving scenario info to DB...")
	var scenarioId int
	db.WithTransaction(func(tx *sqlx.Tx) {
		// TODO: actually capture the user id for the request
		scenarioId = services.CreateScenario(tx, 0, scenarioRequest.Scenario)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			services.CreateDebts(tx, scenarioId, scenarioRequest.Scenario.Debts)
		}()

		go func() {
			defer wg.Done()
			services.CreateSavingsAccounts(tx, scenarioId, scenarioRequest.Scenario.SavingsAccounts)
		}()

		wg.Wait()
	})

	CreateProjections(scenarioRequest)
	return scenarioId
}
