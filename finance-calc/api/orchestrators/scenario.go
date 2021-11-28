package orchestrators

import (
	"log"
	"sync"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/db"
	"financeCalc/api/models"
	"financeCalc/api/services"
	scenarioModels "financeCalc/pkg/scenario/models"
)

func CreateScenario(userId int, scenarioRequest models.ScenarioRequest) int {
	log.Println("Saving scenario info to DB...")
	var scenarioId int
	db.WithTransaction(func(tx *sqlx.Tx) {
		scenarioId = services.CreateScenario(tx, userId, scenarioRequest.Scenario)

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

func GetScenarios(userId int) []*scenarioModels.Scenario {
	log.Printf("Getting scenarios for User %d...\n", userId)
	var scenarios []*scenarioModels.Scenario
	db.WithTransaction(func(tx *sqlx.Tx) {
		scenarios = services.GetScenarios(tx, userId)
	})
	return scenarios
}
