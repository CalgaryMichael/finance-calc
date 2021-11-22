package db

import (
	"github.com/jmoiron/sqlx"

	"financeCalc/api/utils"
)

type TxOperator func(tx *sqlx.Tx)

func WithTransaction(op TxOperator) {
	conn := GetDbInstance()
	tx, err := conn.Beginx()
	utils.CheckError(err)

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()

			// re-raise panic so that the api can return an appropriate response
			panic(err)
		} else {
			commitErr := tx.Commit()
			utils.CheckError(commitErr)
		}
	}()

	op(tx)
}
