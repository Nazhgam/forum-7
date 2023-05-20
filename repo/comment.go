package repo

import (
	"forum/entity"
)

type IComment interface {
	CreateComment(c *entity.Comment) error
	GetCommentByPostID(id int64) ([]entity.Comment, error)
	UpdateComment(c entity.Comment) error
	DeleteCommentByPostID(id int64) error
	DeleteCommentByID(id int64) error
}

func (r repo) CreateComment(c *entity.Comment) error {
	stmt, err := r.db.Prepare("INSERT INTO comments (user_id, post_id, content, created_at, likes, dislikes) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		r.log.Printf("error while to prepare datas to write into the comment table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.UserId, c.PostId, c.Content, c.CreatedAt, c.Likes, c.Dislikes)
	if err != nil {
		r.log.Printf("error while exec prepared datas to write into comment table: %s\n", err.Error())
		return err
	}
	return nil
}

// getCommentByID retrieves a comment from the Comment table by ID
func (r repo) GetCommentByPostID(id int64) ([]entity.Comment, error) {
	query := ` SELECT c.id, c.user_id, c.post_id, c.content, c.created_at,
	   COUNT(CASE WHEN e.likes = 1 THEN e.id END) AS likes_count,
	   COUNT(CASE WHEN e.dislikes = 1 THEN e.id END) AS dislikes_count
	   FROM comments c
	   LEFT JOIN emotions e ON c.id = e.comment_id
	   WHERE c.post_id = ?
	   GROUP BY c.id, c.user_id, c.post_id, c.content, c.created_at;
	`
	rows, err := r.db.Query(query, id)
	if err != nil {
		r.log.Printf("error while to query Get Top Posts By Likes: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	comments := []entity.Comment{}
	for rows.Next() {
		var comment entity.Comment
		err := rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.Likes, &comment.Dislikes)
		if err != nil {
			r.log.Printf("error while to scan Get Comments by Post Id: %s\n", err.Error())
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

// updateComment updates an existing comment in the Comment table
func (r repo) UpdateComment(c entity.Comment) error {
	stmt, err := r.db.Prepare("UPDATE comments SET user_id = ?, post_id = ?, content = ?, updated_at = ?, likes = ?, dislikes = ? WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare update datas in comment table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.UserId, c.PostId, c.Content, c.UpdatedAt, c.Likes, c.Dislikes, c.Id)
	if err != nil {
		r.log.Printf("error while exec prepared update datas in comment table: %s\n", err.Error())
		return err
	}

	return nil
}

// deleteCommentByID deletes a comment from the Comment table by ID
func (r repo) DeleteCommentByPostID(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM comments WHERE post_id = ?")
	if err != nil {
		r.log.Printf("error while to prepare delete comment by id in comment table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		r.log.Printf("error while exec prepared delete comment by id in comment table: %s\n", err.Error())
		return err
	}

	return nil
}

// deleteCommentByID deletes a comment from the Comment table by ID
func (r repo) DeleteCommentByID(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM comments WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare delete comment by id in comment table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		r.log.Printf("error while exec prepared delete comment by id in comment table: %s\n", err.Error())
		return err
	}

	return nil
}
