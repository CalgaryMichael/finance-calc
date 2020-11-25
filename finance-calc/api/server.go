package api

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

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

	log.Fatal(http.ListenAndServe(":3000", r))
}
