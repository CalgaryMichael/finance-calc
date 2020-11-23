package api

import (
	"encoding/json"
	"net/http"

	scenario "financeCalc/pkg/scenario"
	scenarioModels "financeCalc/pkg/scenario/models"
)

func projectScenario(w http.ResponseWriter, req *http.Request) {
	s := scenarioModels.NewFromJSON(req.Body)
	projections := scenario.BuildProjections(s, "DebtTotal", false)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projections)
	return
}
