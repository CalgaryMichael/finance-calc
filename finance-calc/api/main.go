package api

import (
	"io"
	"log"
	http "net/http"

	mux "github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", root)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}
