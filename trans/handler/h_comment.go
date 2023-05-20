package handler

import (
	"encoding/json"
	"forum/entity"
	"net/http"
	"strconv"
	"time"
)

func (h Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	comment := entity.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}

	comment.CreatedAt = time.Now()

	cookie, _ := r.Cookie("Session")
	user := entity.SessionMap[cookie.Value]
	comment.UserId = user.Id

	if err := h.svc.CreateComment(&comment); err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}

	h.Json(w, http.StatusOK, "comment kostym")
	return
}

func (h Handler) DeleteCommentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}
	if err := h.svc.DeleteCommentByID(int64(id)); err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}
	h.Json(w, http.StatusOK, "bzr jok")
	return
}
