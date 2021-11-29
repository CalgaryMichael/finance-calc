package controllers

import (
	"net/http"

	"financeCalc/api/models"
	"financeCalc/api/orchestrators"
	"financeCalc/api/security"
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

func Login(w http.ResponseWriter, req *http.Request) {
	var loginRequest models.LoginRequest
	utils.BindJSON(req.Body, &loginRequest)

	user := orchestrators.GetUserFromCredentials(loginRequest.Email, loginRequest.Password)
	token := security.GenerateUserToken(user)
	resp := models.LoginResponse{
		Token: token,
	}

	utils.JSONResponse(w, 200, resp)
	return
}
