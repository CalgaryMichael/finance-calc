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

func StartServer() {
	log.Println("Starting server on :3000...")
	r := mux.NewRouter()
	registerRoutes(r)

	cors := handlers.CORS(
		handlers.AllowedHeaders(allowedHeaders),
		handlers.AllowedMethods(allowedMethods),
		handlers.AllowedOrigins([]string{"*"}),
	)
	log.Fatal(http.ListenAndServe(":3000", cors(r)))
}
