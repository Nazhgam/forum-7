package middleware

import (
	"context"
	"forum/entity"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func VerifyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Session")
		if err == http.ErrNoCookie {
			cookie = setCookie(w)
			r.AddCookie(cookie)
			switch r.URL.Path {
			case "/home":
				r.URL.Path = "/home"
				r.Method = http.MethodGet
			case "/signup":
				r.URL.Path = "/signupsss"
				r.Method = http.MethodPost
			case "/login":
				r.URL.Path = "/login"
				r.Method = http.MethodPost
			default:
				r.URL.Path = "/home"
				r.Method = http.MethodGet
			}
			next.ServeHTTP(w, r)

			return
		}
		_, ok := entity.SessionMap[cookie.Value]
		if !ok {
			switch r.URL.Path {
			case "/home":
				r.URL.Path = "/home"
				r.Method = http.MethodGet
			case "/signup":
				r.URL.Path = "/signup"
				r.Method = http.MethodPost
			case "/login":
				r.URL.Path = "/login"
				r.Method = http.MethodPost
			default:
				r.URL.Path = "/home"
				r.Method = http.MethodGet
			}

			next.ServeHTTP(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), entity.LoggedIn, true)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func setCookie(w http.ResponseWriter) *http.Cookie {
	uuidString := uuid.New().String()

	cookie := &http.Cookie{
		Name:    entity.Session,
		Value:   uuidString,
		Expires: time.Now().Add(time.Second * 300),
	}

	http.SetCookie(w, cookie)
	return cookie
}
