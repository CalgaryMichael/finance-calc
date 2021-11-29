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
		VALUES (:firstName, :lastName, :email, digest(:password, 'sha256'))
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

func GetUserFromCredentials(tx *sqlx.Tx, email string, password string) (*models.User, error) {
	statement := `
		SELECT
			u.id,
			u.first_name,
			u.last_name,
			u.email
		FROM auth."user" u
		WHERE u.email = :email
			AND u.password = digest(:password, 'sha256')
	`
	params := Params{
		"email":    email,
		"password": password,
	}
	rows, err := tx.NamedQuery(statement, params)
	utils.CheckError(err)

	var user models.User
	rowCount := 0
	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
		)
		utils.CheckError(err)
		rowCount++
		break
	}

	if rowCount < 1 {
		return nil, errors.New(fmt.Sprintf("Unable to find user for %s", email))
	}

	return &user, nil
}
