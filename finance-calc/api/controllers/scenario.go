package controllers

import (
	"io"
	"net/http"

	"financeCalc/api/models"
	"financeCalc/api/utils"
	"financeCalc/pkg/scenario"
)

func NewScenarioRequest(r io.ReadCloser) models.ScenarioRequest {
	var scenarioRequest models.ScenarioRequest
	utils.BindJSON(r, &scenarioRequest)
	return scenarioRequest
}

func CreateScenario(w http.ResponseWriter, req *http.Request) {
	scenarioRequest := NewScenarioRequest(req.Body)
	projections := scenario.BuildProjections(
		scenarioRequest.Scenario,
		scenarioRequest.SortKey,
		scenarioRequest.ReverseSort,
	)

	resp := models.ScenarioResponse{
		Projections: projections,
	}

	utils.JSONResponse(w, 200, resp)
	return
}
