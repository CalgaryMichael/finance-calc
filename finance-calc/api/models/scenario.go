package models

import (
	scenarioModels "financeCalc/pkg/scenario/models"
)

type ScenarioRequest struct {
	Scenario scenarioModels.Scenario `json:"scenario"`
}

type ScenarioResponse struct {
	Projections []*scenarioModels.Projection `json:"projections"`
}
