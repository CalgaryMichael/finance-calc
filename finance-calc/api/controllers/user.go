package controllers

import (
	"net/http"

	"financeCalc/api/models"
	"financeCalc/api/orchestrators"
	"financeCalc/api/utils"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	var createUserRequest models.CreateUserRequest
	utils.BindJSON(req.Body, &createUserRequest)

	userId := orchestrators.CreateUser(createUserRequest.User)
	resp := models.CreateUserResponse{
		UserId: userId,
	}

	utils.JSONResponse(w, 200, resp)
	return
}
