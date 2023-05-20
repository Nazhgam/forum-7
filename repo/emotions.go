package repo

import (
	"fmt"
	"forum/entity"
	"strings"
)

// implement me

// add emotions
// update emotions
// getBYPostID []emotions
// check by post and user id for exist use COUnt func

type IEmotion interface {
	AddEmotion(e *entity.Emotion) error
	UpdateEmotion(e entity.Emotion) error
	GetByEmotionPostId(id int) ([]entity.Emotion, error)
	CheckEmotionForPost(postID, userID int) (bool, error)
	CheckEmotionForComment(commentID, userID int) (bool, error)
	DeleteByPostId(id int) error
	DeleteByCommentId(id int64) error
	DeleteByComments(id []int) error
}

func (r repo) AddEmotion(e *entity.Emotion) error {
	stmt, err := r.db.Prepare("INSERT INTO emotions (post_id, comment_id, user_id, likes, dislikes) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		r.log.Printf("error while to prepare datas to write into the emotion table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.PostID, e.CommentID, e.UserID, e.Likes, e.Dislikes)
	if err != nil {
		r.log.Printf("error while exec prepared datas to write into emotion table: %s\n", err.Error())
		return err
	}
	fmt.Println("norm bd kosildi")
	return nil
}

func (r repo) UpdateEmotion(e entity.Emotion) error {
	stmt, err := r.db.Prepare("UPDATE emotions SET like = ?, dislike = ? WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare update datas in emotion table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Likes, e.Dislikes, e.Id)
	if err != nil {
		r.log.Printf("error while exec prepared update datas in emotion table: %s\n", err.Error())
		return err
	}

	return nil
}

func (r repo) GetByEmotionPostId(id int) ([]entity.Emotion, error) {
	selectQuery := `
		SELECT id, post_id, user_id, like, dislike
		FROM emotions WHERE id = ?
	`
	rows, err := r.db.Query(selectQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	emotions := []entity.Emotion{}
	for rows.Next() {
		var e entity.Emotion
		err := rows.Scan(&e.Id, &e.PostID, &e.UserID, &e.Likes, &e.Dislikes)
		if err != nil {
			return nil, err
		}
		emotions = append(emotions, e)
	}

	return emotions, nil
}

func (r repo) CheckEmotionForPost(postID, userID int) (bool, error) {
	fmt.Println("income ", postID, userID)
	countQuery := `
		SELECT COUNT(*)
		FROM emotions
		WHERE post_id = ? AND user_id = ?
	`
	var count int
	err := r.db.QueryRow(countQuery, postID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	fmt.Println(count)
	return count > 0, nil
}

func (r repo) CheckEmotionForComment(commentID, userID int) (bool, error) {
	countQuery := `
		SELECT COUNT(*)
		FROM emotions
		WHERE comment_id = ? AND user_id = ?
	`
	var count int
	err := r.db.QueryRow(countQuery, commentID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r repo) DeleteByPostId(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM emotions WHERE post_id = ?")
	if err != nil {
		r.log.Printf("error while to prepare delete user by id in post table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		r.log.Printf("error while exec prepared delete user by id in post table: %s\n", err.Error())
		return err
	}

	return nil
}

func (r repo) DeleteByComments(id []int) error {
	placeHolder := make([]string, len(id))
	for i := range placeHolder {
		placeHolder[i] = fmt.Sprintf("%d", id[i])
	}
	strHold := strings.Join(placeHolder, ",")

	stmt, err := r.db.Prepare(fmt.Sprintf("DELETE FROM emotions WHERE comment_id IN (%s)", strHold))
	if err != nil {
		r.log.Printf("error while to prepare delete user by id in post table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		r.log.Printf("error while exec prepared delete user by id in post table: %s\n", err.Error())
		return err
	}

	return nil
}

func (r repo) DeleteByCommentId(id int64) error {
	stmt, err := r.db.Prepare(fmt.Sprintf("DELETE FROM emotions WHERE comment_id=?"))
	if err != nil {
		r.log.Printf("error while to prepare delete user by id in post table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		r.log.Printf("error while exec prepared delete user by id in post table: %s\n", err.Error())
		return err
	}

	return nil
}
