package daos

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"financeCalc/api/models"
	"financeCalc/api/utils"
)

func CreateUser(tx *sqlx.Tx, user models.User) int {
	statement := `
		INSERT INTO auth."user" (first_name, last_name, email, password)
		VALUES (:firstName, :lastName, :email, :password)
		RETURNING id;
	`
	params := Params{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"password":  user.Password,
	}
	rows, err := tx.NamedQuery(statement, params)
	utils.CheckError(err)

	id, ok := GetInsertedId(rows)
	if !ok {
		panic(errors.New(fmt.Sprintf("Unable to insert User \"%d\"", user.Email)))
	}
	return id
}
