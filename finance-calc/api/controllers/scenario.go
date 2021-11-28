package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"financeCalc/api/models"
	"financeCalc/api/orchestrators"
	"financeCalc/api/utils"
)

func CreateScenario(w http.ResponseWriter, req *http.Request) {
	var scenarioRequest models.ScenarioRequest
	utils.BindJSON(req.Body, &scenarioRequest)

	// TODO: actually capture the user id for the request
	userId := 1

	scenarioId := orchestrators.CreateScenario(userId, scenarioRequest)
	resp := models.ScenarioResponse{
		ScenarioId: scenarioId,
	}

	utils.JSONResponse(w, 200, resp)
	return
}

func GetScenarios(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	param, present := query["userId"]
	if !present {
		panic(errors.New("Please provide User ID"))
	}
	userId, err := strconv.Atoi(param[0])
	utils.CheckError(err)

	scenarios := orchestrators.GetScenarios(userId)
	resp := models.GetScenariosResponse{
		Scenarios: scenarios,
	}

	utils.JSONResponse(w, 200, resp)
	return
}
