package entity

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	SessionMap  = make(map[string]User)
)

const (
	Session           = "Session"
	LoggedIn loggedIn = "LoggedIn"
)

var Categories = map[string]bool{
	"Anime":       true,
	"Sport":       true,
	"Cars":        true,
	"Education":   true,
	"Cyber Sport": true,
	"Books":       true,
	"Camping":     true,
	"Movies":      true,
}

type loggedIn string

type User struct {
	Id        int64      `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Post struct {
	Id        int64      `json:"id"`
	Category  []string   `json:"category"`
	UserId    int64      `json:"user_id"`
	UserName  string     `json:"user_name"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Likes     int64      `json:"likes"`
	Dislikes  int64      `json:"dislikes"`
}

type Category struct {
	Id     int64  `json:"id"`
	PostId int64  `json:"post_id"`
	Name   string `json:"name"`
}

type Comment struct {
	Id        int64      `json:"id"`
	UserId    int64      `json:"user_id"`
	PostId    int64      `json:"post_id"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Likes     int64      `json:"likes"`
	Dislikes  int64      `json:"dislikes"`
}

type Profile struct {
	Id        int64      `json:"id"`
	UserId    int64      `json:"user_id"`
	Name      string     `json:"name"`
	Bio       string     `json:"bio"`
	ImageUrl  string     `json:"image_url"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Emotion struct {
	Id        int64 `json:"id"`
	PostID    int64 `json:"post_id"`
	CommentID int64 `json:"comment_id"`
	UserID    int64 `json:"user_id"`
	Likes     bool  `json:"likes"`
	Dislikes  bool  `json:"dislikes"`
}

type PostByID struct {
	Post
	Comments []Comment
}
