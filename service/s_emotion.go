package service

import (
	"errors"
	"fmt"
	"forum/entity"
)

func (s *service) AddEmotion(e *entity.Emotion) error {
	if (e.Likes && e.Dislikes) || (!e.Likes && !e.Dislikes) {
		return errors.New("Error emotion service (likes and dislikes)")
	}

	if e.CommentID == 0 && e.PostID == 0 {
		return errors.New("Error emotion service (post_id and comment_id)")
	}

	check, err := s.CheckEmotion(e)
	if err != nil {
		return err
	}
	switch check {
	case 0:
		return s.repo.AddEmotion(e)
	case 1:
		return s.repo.UpdateEmotion(*e)
	default:
		return s.repo.DeleteEmotionById(int(e.Id))
	}
}

func (s *service) CheckEmotion(e *entity.Emotion) (int, error) {
	res, err := s.repo.GetEmotionByPostCommentId(int(e.PostID), int(e.CommentID))
	if err != nil {
		return 0, err
	}
	var emotions []entity.Emotion
	for i := range res {
		if res[i].UserID == e.UserID {
			emotions = append(emotions, res[i])
		}
	}
	fmt.Println(emotions)
	switch len(emotions) {
	case 0:
		return 0, nil
	case 1:
		e.Id = emotions[0].Id
		if e.Likes == emotions[0].Likes && e.Dislikes == emotions[0].Dislikes {
			return 2, nil
		} else {
			return 1, nil
		}
	default:
		return 0, errors.New("found a lot of emotion by given parameters")
	}
}
