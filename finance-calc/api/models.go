package api

import (
	"encoding/json"
	"io"

	scenarioModels "financeCalc/pkg/scenario/models"
)

type ScenarioRequest struct {
	Scenario    scenarioModels.Scenario `json:"scenario"`
	SortKey     string                  `json:"sortKey"`
	ReverseSort bool                    `json:"reverseSort"`
}

func NewScenarioRequest(r io.ReadCloser) ScenarioRequest {
	decoder := json.NewDecoder(r)
	var scenarioRequest ScenarioRequest
	err := decoder.Decode(&scenarioRequest)
	if err != nil {
		panic(err)
	}
	return scenarioRequest
}

type ScenarioResponse struct {
	Projections []*scenarioModels.Projection `json:"projections"`
}
