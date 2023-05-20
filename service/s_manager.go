package service

import (
	"forum/entity"
	"forum/repo"
	"log"
)

type IService interface {
	GetMostLikedCategoryPosts(category string) ([]entity.Post, error)
	GetMostLikedPosts() ([]entity.Post, error)

	GetAllPosts() ([]entity.Post, error)
	GetPostByID(id int) (*entity.Post, error)

	SignUp(user *entity.User) error
	LogIn(user *entity.User) (entity.User, error)

	CreatePost(post *entity.Post) error
	Update(p entity.Post) (*entity.Post, error)

	CreateComment(comment *entity.Comment) error
	GetCommentByPostID(id int64) ([]entity.Comment, error)
	DeleteComments(postId int64) error
	DeleteCommentByID(id int64) error

	AddEmotionToPost(e *entity.Emotion) error
	AddEmotionToComment(e *entity.Emotion) error
}

type service struct {
	log  *log.Logger
	repo repo.IDb
}

func New(repo repo.IDb, l *log.Logger) IService {
	return &service{log: l, repo: repo}
}
