package api

import (
	"encoding/json"
	"net/http"
)

type jsonOK struct {
	Data interface{} `json:"data"`
}

type jsonErr struct {
	Err string `json:"error"`
}

func sendOK(w http.ResponseWriter, status int, payload interface{}) {
	resp := jsonOK{
		Data: payload,
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func sendERR(w http.ResponseWriter, status int, err error) {
	resp := jsonErr{
		Err: err.Error(),
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
