package tools

import (
	"encoding/json"
	"net/http"
)

// Response monta uma resposta de devolução a chamada à API
func (t *Tools) Response(w http.ResponseWriter, status int, body interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if status >= 500 {
		w.Write([]byte(body.(string)))
		return
	}

	if body != nil {
		rJSON, _ := json.Marshal(body)
		w.Write(rJSON)
	}
}
