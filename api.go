package main

import (
	"encoding/json"
	"net/http"
)

func in(w http.ResponseWriter, r *http.Request, input interface{}) bool {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(input); err != nil {
		nok(w, 400, err.Error())
		return false
	}
	return true
}

func ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	var output struct {
		Success bool        `json:"success"`
		Result  interface{} `json:"result"`
	}
	output.Success = true
	output.Result = data

	enc := json.NewEncoder(w)
	enc.Encode(output)
}

func nok(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	var output struct {
		Success bool        `json:"success"`
		Error   interface{} `json:"error"`
	}
	output.Error = data

	enc := json.NewEncoder(w)
	enc.Encode(output)
}
