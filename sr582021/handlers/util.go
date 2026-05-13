package handlers

import (
	"encoding/json"
	"net/http"
)

func renderJSON(w http.ResponseWriter, body interface{}, statusCode int){
	if body == nil {
		w.WriteHeader(statusCode)
		return
	}
	json, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}