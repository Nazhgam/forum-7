package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"forum/entity"
	"forum/entity/cerror"
)

func (h Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := entity.Post{}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	post.CreatedAt = time.Now()

	cookie, _ := r.Cookie("Session")
	user := entity.SessionMap[cookie.Value]
	post.UserId = user.Id

	if err := h.svc.CreatePost(&post); err != nil {
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

func (h Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := entity.Post{}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
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

	resPost, err := h.svc.GetPostByID(id)
	if err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return

	}
	if resPost == nil {
		h.Json(w, http.StatusNotFound, cerror.NewError("Brat ondai zapis jok"))
		return
	}

	resCom, err := h.svc.GetCommentByPostID(int64(id))
	if err != nil {
		h.Json(w, http.StatusInternalServerError, err)
		return
	}

	response := struct {
		Post    *entity.Post
		Comment []entity.Comment
	}{
		Post:    resPost,
		Comment: resCom,
	}

	h.Json(w, http.StatusOK, response)
	return
}

func (h Handler) MostLikedCategory(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.GetMostLikedCategoryPosts("Anime")
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}

func (h Handler) MostLikedPost(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.GetMostLikedPosts()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}
