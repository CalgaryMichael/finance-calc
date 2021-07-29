package api

import (
	"log"
	"net/http"
)

type MuxRouteFunc func(http.ResponseWriter, *http.Request)

func requestLogger(f MuxRouteFunc) MuxRouteFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v\n", r.Method, r.URL)
		f(w, r)
	}
}
