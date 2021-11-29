package services

import (
	"github.com/jmoiron/sqlx"

	"financeCalc/api/daos"
	"financeCalc/api/models"
	"financeCalc/api/utils"
)

func CreateUser(tx *sqlx.Tx, user models.User) int {
	return daos.CreateUser(tx, user)
}

func GetUserFromCredentials(tx *sqlx.Tx, email string, password string) *models.User {
	user, err := daos.GetUserFromCredentials(tx, email, password)
	utils.CheckError(err)
	return user
}
