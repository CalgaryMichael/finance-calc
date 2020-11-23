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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projections)
	return
}
