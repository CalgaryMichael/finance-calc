package models

import (
	scenarioModels "financeCalc/pkg/scenario/models"
)

type ProjectionRequest struct {
	ScenarioId int `json:"scenarioId"`
}

type ProjectionResponse struct {
	ScenarioId  int                          `json:"scenarioId"`
	Projections []*scenarioModels.Projection `json:"projections"`
}
