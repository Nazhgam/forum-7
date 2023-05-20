package handler

import (
	"encoding/json"
	"forum/entity"
	"forum/entity/cerror"
	"net/http"
	"strconv"
	"time"
)

func (h Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := entity.Post{}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}

	post.CreatedAt = time.Now()

	cookie, _ := r.Cookie("Session")
	user := entity.SessionMap[cookie.Value]
	post.UserId = user.Id

	if err := h.svc.CreatePost(&post); err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}
	h.Json(w, http.StatusOK, "kostim")
	return
}

func (h Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := entity.Post{}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.Update(post)
	if err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}
	h.Json(w, http.StatusOK, res)
}

func (h Handler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		h.Json(w, http.StatusBadRequest, err)
		return
	}

	resPost, err := h.svc.GetPostByID(int64(id))
	if err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return

	}
	if resPost == nil {
		h.Json(w, http.StatusNotFound, cerror.NewError("Brat ondai zapis jok"))
		return
	}

	h.Json(w, http.StatusOK, resPost)
	return
}

func (h Handler) MostLikedCategory(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.GetMostLikedCategoryPosts("Anime")
	if err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

func (h Handler) MostLikedPost(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.GetMostLikedPosts()
	if err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
	return
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.Json(w, http.StatusBadRequest, err)
		return
	}
	if err := h.svc.DeletePostByID(int64(id)); err != nil {
		h.Json(w, http.DefaultMaxHeaderBytes, err)
		return
	}
	h.Json(w, http.StatusOK, "BAZAR JOK")
}
