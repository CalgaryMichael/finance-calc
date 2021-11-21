package controllers

import (
	"io"
	"net/http"

	"financeCalc/api/models"
	"financeCalc/api/orchestrators"
	"financeCalc/api/utils"
)

func NewScenarioRequest(r io.ReadCloser) models.ScenarioRequest {
	var scenarioRequest models.ScenarioRequest
	utils.BindJSON(r, &scenarioRequest)
	return scenarioRequest
}

func CreateScenario(w http.ResponseWriter, req *http.Request) {
	scenarioRequest := NewScenarioRequest(req.Body)
	projections := orchestrators.CreateScenario(scenarioRequest)

	resp := models.ScenarioResponse{
		Projections: projections,
	}

	utils.JSONResponse(w, 200, resp)
	return
}
