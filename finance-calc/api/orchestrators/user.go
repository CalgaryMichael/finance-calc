package orchestrators

import (
	"log"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/db"
	"financeCalc/api/models"
	"financeCalc/api/services"
)

func CreateUser(user models.User) int {
	log.Printf("Creating user for email %s...\n", user.Email)
	var userId int
	db.WithTransaction(func(tx *sqlx.Tx) {
		userId = services.CreateUser(tx, user)
	})
	return userId
}

func GetUserFromCredentials(email string, password string) *models.User {
	log.Printf("Getting user for email %s...\n", email)
	var user *models.User
	db.WithTransaction(func(tx *sqlx.Tx) {
		user = services.GetUserFromCredentials(tx, email, password)
	})
	return user
}
