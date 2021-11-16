package main

import (
	db "financeCalc/api/db"
	api "financeCalc/api/server"
)

func main() {
	db.Connect()
	api.StartServer()
}
