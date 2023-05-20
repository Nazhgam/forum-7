package service

import (
	"errors"
	"forum/entity"
)

func (s *service) CreateComment(comment *entity.Comment) error {
	if len(comment.Content) < 1 {
		return errors.New("comment content empty")
	}
	if err := s.repo.CreateComment(comment); err != nil {
		return err
	}

	return nil
}

func (s *service) GetCommentByPostID(id int64) ([]entity.Comment, error) {
	return s.repo.GetCommentByPostID(id)
}

func (s *service) DeleteCommentByID(id int64) error {
	if err := s.repo.DeleteByCommentId(id); err != nil {
		return err
	}
	return s.repo.DeleteCommentByID(id)
}

func (s *service) DeleteComments(postId int64) error {
	comments, err := s.repo.GetCommentByPostID(postId)
	if err != nil {
		return err
	}
	ids := make([]int, len(comments))
	for i := range comments {
		ids[i] = int(comments[i].Id)
	}
	if err := s.repo.DeleteByComments(ids); err != nil {
		return err
	}
	return s.repo.DeleteCommentByPostID(postId)
}
