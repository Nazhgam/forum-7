package handler

import (
	"encoding/json"
	"net/http"
)

func (h Handler) Json(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	switch payload.(type) {
	case error:
		w.Write([]byte(payload.(error).Error()))
	default:
		body, err := json.Marshal(payload)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write(body)
	}

}
