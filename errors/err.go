package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleErr(w http.ResponseWriter, err error) {
	res := err.Error()
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Printf("error sending response %v\n", err)
	}

	w.WriteHeader(http.StatusBadRequest)
}

func HandleDBFromContextErr(w http.ResponseWriter) {
	res := "error"
	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Printf("error sending response %v\n", err)
	}

	w.WriteHeader(http.StatusBadRequest)
}
