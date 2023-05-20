package handler

import (
	"encoding/json"
	"net/http"

	"forum/entity"
)

func (h Handler) AddEmotionToPost(w http.ResponseWriter, r *http.Request) {
	emotion := entity.Emotion{}

	if err := json.NewDecoder(r.Body).Decode(&emotion); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	cookie, _ := r.Cookie("Session")
	user := entity.SessionMap[cookie.Value]
	emotion.UserID = user.Id

	if err := h.svc.AddEmotionToPost(&emotion); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("kostim"))
	return
}

func (h Handler) AddEmotionToComment(w http.ResponseWriter, r *http.Request) {
	emotion := entity.Emotion{}

	if err := json.NewDecoder(r.Body).Decode(&emotion); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	cookie, _ := r.Cookie("Session")
	user := entity.SessionMap[cookie.Value]
	emotion.UserID = user.Id

	if err := h.svc.AddEmotionToComment(&emotion); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("kostim"))
	return
}
