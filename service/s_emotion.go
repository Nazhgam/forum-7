package service

import (
	"errors"
	"fmt"

	"forum/entity"
)

func (s *service) AddEmotionToPost(e *entity.Emotion) error {
	if (e.Likes && e.Dislikes) || (!e.Likes && !e.Dislikes) {
		return errors.New("Error emotion service (likes and dislikes)")
	}

	if e.PostID == 0 {
		return errors.New("Error emotion service (post_id and comment_id)")
	}
	fmt.Println("service", e)

	exist, err := s.repo.CheckEmotionForPost(int(e.PostID), int(e.UserID))
	if err != nil {
		return err
	}

	if exist {
		fmt.Println("bargoi uje like")
		return nil
	}
	return s.repo.AddEmotion(e)
}

func (s *service) AddEmotionToComment(e *entity.Emotion) error {
	if (e.Likes && e.Dislikes) || (!e.Likes && !e.Dislikes) {
		return errors.New("Error emotion service (likes and dislikes)")
	}

	if e.CommentID == 0 {
		return errors.New("Error emotion service (post_id and comment_id)")
	}

	exist, err := s.repo.CheckEmotionForComment(int(e.CommentID), int(e.UserID))
	if err != nil {
		return err
	}

	if exist {
		return nil
	}
	return s.repo.AddEmotion(e)
}
