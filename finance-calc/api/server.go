package api

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var allowedHeaders []string = []string{
	"X-Requested-With",
	"Content-Type",
	"Authorization",
}

var allowedMethods []string = []string{
	"GET",
	"POST",
	"PUT",
	"HEAD",
	"OPTIONS",
}

type MuxRouteFunc func(http.ResponseWriter, *http.Request)

func requestLogger(f MuxRouteFunc) MuxRouteFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v\n", r.Method, r.URL)
		f(w, r)
	}
}

func StartServer() {
	log.Println("Starting server on :3000...")
	r := mux.NewRouter()
	r.HandleFunc("/project", requestLogger(projectScenario)).Methods("POST")

	cors := handlers.CORS(
		handlers.AllowedHeaders(allowedHeaders),
		handlers.AllowedMethods(allowedMethods),
		handlers.AllowedOrigins([]string{"*"}),
	)
	log.Fatal(http.ListenAndServe(":3000", cors(r)))
}
