package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"financeCalc/api/controllers"
	"financeCalc/api/utils"
)

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/scenario", utils.HandleErrors(controllers.CreateScenario)).Methods(http.MethodPost)
}
