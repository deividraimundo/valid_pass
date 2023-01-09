package tools

import (
	"encoding/json"
	"net/http"
)

func (t *Tools) Response(w http.ResponseWriter, status int, body interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if body != nil {
		rJSON, _ := json.Marshal(body)
		w.Write(rJSON)
	}
}
