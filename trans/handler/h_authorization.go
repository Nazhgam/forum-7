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
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	req.CreatedAt = time.Now()
	h.log.Println("start svs signup")
	if err := h.svc.SignUp(&req); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(err.Error()))
		return
	}

	cookie, err := r.Cookie("Session")
	if err != nil || err == http.ErrNoCookie {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	entity.SessionMap[cookie.Value] = req
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("kostim"))
	return
}

func (h Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	req := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := h.svc.LogIn(&req)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	cookie, err := r.Cookie("Session")
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var mu sync.Mutex
	mu.Lock()
	entity.SessionMap[cookie.Value] = user
	mu.Unlock()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Yeaaaaah hello baby"))
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

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
