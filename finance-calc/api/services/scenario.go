package services

import (
	"financeCalc/api/daos"
	scenarioModels "financeCalc/pkg/scenario/models"

	"github.com/jmoiron/sqlx"
)

func CreateScenario(tx *sqlx.Tx, userId int, scenario scenarioModels.Scenario) int {
	scenario.Id = daos.CreateScenario(tx, userId, scenario)
	return scenario.Id
}
