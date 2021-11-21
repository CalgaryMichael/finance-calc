package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"financeCalc/api/controllers"
)

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/scenario", controllers.CreateScenario).Methods(http.MethodPost)
}
