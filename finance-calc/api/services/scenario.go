package services

import (
	"financeCalc/api/daos"
	scenarioModels "financeCalc/pkg/scenario/models"

	"github.com/jmoiron/sqlx"
)

func CreateScenario(tx *sqlx.Tx, userId int, scenario scenarioModels.Scenario) int {
	return daos.CreateScenario(tx, userId, scenario)
}
