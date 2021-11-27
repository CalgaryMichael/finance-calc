package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"financeCalc/api/models"
	"financeCalc/api/orchestrators"
	"financeCalc/api/utils"
)

func NewProjectionRequest(req *http.Request) models.ProjectionRequest {
	routeVars := mux.Vars(req)
	scenarioId, err := strconv.Atoi(routeVars["id"])
	utils.CheckError(err)
	return models.ProjectionRequest{
		ScenarioId: scenarioId,
	}
}

func GetProjectionsForScenario(w http.ResponseWriter, req *http.Request) {
	projectionRequest := NewProjectionRequest(req)
	projections := orchestrators.GetProjectionsForScenario(projectionRequest)

	resp := models.ProjectionResponse{
		ScenarioId:  projectionRequest.ScenarioId,
		Projections: projections,
	}

	utils.JSONResponse(w, 200, resp)
	return
}
