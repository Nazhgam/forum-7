package handler

import (
	"encoding/json"
	"forum/entity"
	"net/http"
	"sync"
	"time"
)

func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	h.log.Println("START handler signup")
	req := entity.User{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}

	req.CreatedAt = time.Now()
	h.log.Println("start svs signup")
	if err := h.svc.SignUp(&req); err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}

	cookie, err := r.Cookie("Session")
	if err != nil || err == http.ErrNoCookie {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}

	entity.SessionMap[cookie.Value] = req
	h.Json(w, http.StatusOK, "kostim")
	return
}

func (h Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	req := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.svc.LogIn(&req)
	if err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}
	cookie, err := r.Cookie("Session")
	if err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}
	var mu sync.Mutex
	mu.Lock()
	entity.SessionMap[cookie.Value] = user
	mu.Unlock()
	h.Json(w, http.StatusOK, "Yeaaaah hello baby")
	return
}

func (h Handler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Session")

	delete(entity.SessionMap, cookie.Value)

	cookie = &http.Cookie{
		Name:   "Session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
