package api

import (
	"log"
	http "net/http"

	mux "github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/project", projectScenario).Methods("POST")

	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
	log.Print("Now listening on port 3000...")
}
