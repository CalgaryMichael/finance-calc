package services

import (
	"github.com/jmoiron/sqlx"

	"financeCalc/api/daos"
	"financeCalc/api/models"
)

func CreateUser(tx *sqlx.Tx, user models.User) int {
	return daos.CreateUser(tx, user)
}
