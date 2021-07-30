package utils

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	setStatus(w, status)
	err := setResponseBody(w, resp)
	if err != nil {
		setStatus(w, 500)
	}
	return
}

func setResponseBody(w http.ResponseWriter, resp interface{}) error {
	respBytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = w.Write(respBytes)
	if err != nil {
		return err
	}
	return nil
}

func setStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
