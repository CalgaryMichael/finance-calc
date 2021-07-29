package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"financeCalc/api/models"
	"financeCalc/pkg/scenario"
)

func NewScenarioRequest(r io.ReadCloser) models.ScenarioRequest {
	decoder := json.NewDecoder(r)
	var scenarioRequest models.ScenarioRequest
	err := decoder.Decode(&scenarioRequest)
	if err != nil {
		panic(err)
	}
	return scenarioRequest
}

func CreateScenario(w http.ResponseWriter, req *http.Request) {
	scenarioRequest := NewScenarioRequest(req.Body)
	projections := scenario.BuildProjections(
		scenarioRequest.Scenario,
		scenarioRequest.SortKey,
		scenarioRequest.ReverseSort,
	)

	response := models.ScenarioResponse{
		Projections: projections,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
