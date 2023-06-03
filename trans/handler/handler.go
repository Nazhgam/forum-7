package handler

import (
	"encoding/json"
	"forum/service"
	"log"
	"net/http"
)

type Handler struct {
	log *log.Logger
	svc service.IService
}

func New(svc service.IService, l *log.Logger) Handler {
	return Handler{log: l, svc: svc}
}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.GetAllPosts()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}
