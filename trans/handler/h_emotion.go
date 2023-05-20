package handler

import (
	"encoding/json"
	"forum/entity"
	"net/http"
)

func (h Handler) AddEmotionToPost(w http.ResponseWriter, r *http.Request) {
	emotion := entity.Emotion{}

	if err := json.NewDecoder(r.Body).Decode(&emotion); err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}
	cookie, _ := r.Cookie("Session")
	user := entity.SessionMap[cookie.Value]
	emotion.UserID = user.Id

	if err := h.svc.AddEmotion(&emotion); err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}
	h.Json(w, http.StatusOK, "kostim")
	return
}

func (h Handler) AddEmotionToComment(w http.ResponseWriter, r *http.Request) {
	emotion := entity.Emotion{}

	if err := json.NewDecoder(r.Body).Decode(&emotion); err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}

	cookie, _ := r.Cookie("Session")
	user := entity.SessionMap[cookie.Value]
	emotion.UserID = user.Id

	if err := h.svc.AddEmotion(&emotion); err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}
	h.Json(w, http.StatusOK, "kostim")
	return
}
