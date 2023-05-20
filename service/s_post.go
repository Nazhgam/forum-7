package service

import (
	"errors"
	"forum/entity"
)

func (s *service) GetPostByID(id int64) (p *entity.PostByID, err error) {
	res, err := s.repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	comments, err := s.GetCommentByPostID(id)
	if err != nil {
		return nil, err
	}

	p = &entity.PostByID{
		Post:     *res,
		Comments: comments,
	}

	return p, nil
}

func (s *service) GetAllPosts() ([]entity.Post, error) {
	return s.repo.GetAllPosts()
}

func (s *service) GetMostLikedCategoryPosts(category string) ([]entity.Post, error) {
	return s.repo.GetTopPostsByCategoryLikes(category)
}

func (s *service) GetMostLikedPosts() ([]entity.Post, error) {
	return s.repo.GetTopPostsByLikes()
}

func (s *service) CreatePost(post *entity.Post) error {
	if len(post.Title) < 1 || len(post.Content) < 1 {
		return errors.New("Title or content empty")
	}

	if err := s.repo.CreatePost(post); err != nil {
		return err
	}

	for _, v := range post.Category {
		if entity.Categories[v] == true {
			categ := entity.Category{
				PostId: post.Id,
				Name:   v,
			}
			if err := s.repo.CreateCategory(&categ); err != nil {
				return errors.New("Kate create")
			}
		} else {
			return errors.New("Kate kategoria")
		}
	}

	return nil
}

func (s *service) Update(p entity.Post) (*entity.Post, error) {
	if len(p.Content) == 0 {
		return nil, errors.New("empty content on update")
	}
	return s.repo.UpdatePost(p)
}

func (s *service) DeletePostByID(postId int64) error {
	if err := s.DeleteComments(postId); err != nil {
		return err
	}

	if err := s.repo.DeleteByPostId(int(postId)); err != nil {
		return err
	}
	return s.repo.DeletePostByID(postId)
}
