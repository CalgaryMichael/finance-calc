package api

import (
	"github.com/gorilla/mux"

	"financeCalc/api/controllers"
)

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/scenario", controllers.CreateScenario).Methods("POST")
}
