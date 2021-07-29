package models

import (
	scenarioModels "financeCalc/pkg/scenario/models"
)

type ScenarioRequest struct {
	Scenario    scenarioModels.Scenario `json:"scenario"`
	SortKey     string                  `json:"sortKey"`
	ReverseSort bool                    `json:"reverseSort"`
}

type ScenarioResponse struct {
	Projections []*scenarioModels.Projection `json:"projections"`
}
