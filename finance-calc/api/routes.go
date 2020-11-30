package api

import (
	"encoding/json"
	"net/http"

	scenario "financeCalc/pkg/scenario"
)

func projectScenario(w http.ResponseWriter, req *http.Request) {
	scenarioRequest := NewScenarioRequest(req.Body)
	projections := scenario.BuildProjections(
		scenarioRequest.Scenario,
		scenarioRequest.SortKey,
		scenarioRequest.ReverseSort,
	)

	response := ScenarioResponse{
		Projections: projections,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
