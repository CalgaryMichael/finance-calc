package server

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
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodHead,
	http.MethodOptions,
}

var allowedOrigins []string = []string{
	"http://localhost:8080",
}

func StartServer() {
	log.Println("Starting server on :3000...")
	r := mux.NewRouter()
	registerRoutes(r)

	r.Use(loggingMiddleware)

	cors := handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowedHeaders(allowedHeaders),
		handlers.AllowedMethods(allowedMethods),
	)
	log.Fatal(http.ListenAndServe(":3000", cors(r)))
}
