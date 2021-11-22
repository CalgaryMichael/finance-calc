package utils

import (
	"log"
	"net/http"
)

type ControllerFunc func(w http.ResponseWriter, req *http.Request)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleErrors(op ControllerFunc) ControllerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			// determine if there were any panic's thrown during the function's execution
			if err := recover(); err != nil {
				log.Print(err)
				resp := map[string]string{
					"error": "Unable to complete request",
				}
				JSONResponse(w, 500, resp)
			}
		}()

		// call underlying function
		op(w, req)
	}
}
