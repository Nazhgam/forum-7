package repo

import (
	"forum/entity"
	"time"
)

type IPost interface {
	CreatePost(p *entity.Post) error
	GetPostByID(id int64) (*entity.Post, error)
	UpdatePost(p entity.Post) (*entity.Post, error)
	DeletePostByID(id int64) error
	GetTopPostsByLikes() ([]entity.Post, error)
	GetTopPostsByCategoryLikes(category string) ([]entity.Post, error)
	// GetCommentIds(id int) ([]int, error)
	GetAllPosts() ([]entity.Post, error)
}

func (r repo) CreatePost(p *entity.Post) error {
	stmt, err := r.db.Prepare("INSERT INTO posts (user_id, title, content, created_at) VALUES ($1, $2, $3, $4)")
	if err != nil {
		r.log.Printf("error while to prepare post datas to write into the post table: %s\n", err.Error())
		return err
	}

	defer stmt.Close()

	stm, err := stmt.Exec(p.UserId, p.Title, p.Content, time.Now())
	if err != nil {
		r.log.Printf("error while exec prepared post datas to write into post table: %s\n", err.Error())
		return err
	}

	p.Id, err = stm.LastInsertId()
	if err != nil {
		r.log.Printf("error while exec prepared datas to write into post table: %s\n", err.Error())
		return err
	}

	return nil
}

// getPostByID retrieves a post from the Post table by ID
func (r repo) GetPostByID(id int64) (*entity.Post, error) {
	rows, err := r.db.Query(`
	SELECT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, c.name, 
    COUNT(CASE WHEN e.likes = 1 THEN e.id END) AS likes_count,
    COUNT(CASE WHEN e.dislikes = 1 THEN e.id END) AS dislikes_count
	FROM posts p
	LEFT JOIN category c ON p.id = c.post_id
	LEFT JOIN emotions e ON p.id = e.post_id
	WHERE p.id=?
    GROUP BY p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, c.name
	ORDER BY p.created_at DESC
	`, id)
	if err != nil {
		r.log.Printf("error while to prepare datas to get post by id from post table: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var post entity.Post

	for rows.Next() {
		var cur entity.Post
		categ := ""
		err := rows.Scan(&cur.Id, &cur.Title, &cur.Content, &cur.UserId, &cur.CreatedAt, &cur.UpdatedAt, &categ, &cur.Likes, &cur.Dislikes)
		if err != nil {
			r.log.Printf("error while to scan Get Posts By Category Or Title: %s\n", err.Error())
			return nil, err
		}
		if post.Id == 0 {
			post = cur
		}
		post.Category = append(post.Category, categ)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	if post.Id == 0 {
		return nil, nil
	}

	return &post, nil
}

// updatePost updates an existing post in the Post table
func (r repo) UpdatePost(p entity.Post) (*entity.Post, error) {
	stmt, err := r.db.Prepare("UPDATE posts SET content = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare update datas in post table: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Content, time.Now(), p.Id)
	if err != nil {
		r.log.Printf("error while exec prepared update datas in post table: %s\n", err.Error())
		return nil, err
	}

	return r.GetPostByID(p.Id)
}

// deletePostByID deletes a post from the Post table by ID
func (r repo) DeletePostByID(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM posts WHERE id = ?")
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

// получить 10 постов из базы данных с наибольшим количеством лайков
func (r repo) GetTopPostsByCategoryLikes(category string) ([]entity.Post, error) {
	query := `
	SELECT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at,
    COUNT(CASE WHEN e.likes = 1 THEN e.id END) AS likes_count,
    COUNT(CASE WHEN e.dislikes = 1 THEN e.id END) AS dislikes_count
	FROM posts p
	LEFT JOIN category c ON p.id = c.post_id
	LEFT JOIN emotions e ON p.id = e.post_id
	WHERE c.name=? 
    GROUP BY p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, c.name
	ORDER BY likes_count DESC
	`
	rows, err := r.db.Query(query, category)
	if err != nil {
		r.log.Printf("error while to query Get Top Posts By Likes: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	posts := []entity.Post{}
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.Dislikes)
		if err != nil {
			r.log.Printf("error while to scan Get Top Posts By Likes: %s\n", err.Error())
			return nil, err
		}
		post.Category = append(post.Category, category)
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		r.log.Println(err)
		return nil, err
	}

	return posts, nil
}

func (r repo) GetCommentIds(id int) ([]int, error) {
	rows, err := r.db.Query(`SELECT comment_id from posts where id=?`, id)
	if err != nil {
		return nil, err
	}
	var res []int
	for rows.Next() {
		curId := 0
		if err := rows.Scan(&curId); err != nil {
			return nil, err
		}
		res = append(res, curId)
	}
	return res, nil
}

func (r repo) GetTopPostsByLikes() ([]entity.Post, error) {
	query := `
	SELECT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at,
    COUNT(CASE WHEN e.likes = 1 THEN e.id END) AS likes_count,
    COUNT(CASE WHEN e.dislikes = 1 THEN e.id END) AS dislikes_count
	FROM posts p
	LEFT JOIN category c ON p.id = c.post_id
	LEFT JOIN emotions e ON p.id = e.post_id
	
    GROUP BY p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, c.name
	ORDER BY  p.created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		r.log.Printf("error while to query Get Top Posts By Likes: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	posts := []entity.Post{}
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.Dislikes)
		if err != nil {
			r.log.Printf("error while to scan Get Top Posts By Likes: %s\n", err.Error())
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		r.log.Println(err)
		return nil, err
	}

	return posts, nil
}

// получить все посты из базы данных в порядке, определенном запросом
func (r repo) GetAllPosts() ([]entity.Post, error) {
	rows, err := r.db.Query(`
	SELECT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, c.name, 
    COUNT(CASE WHEN e.likes = 1 THEN e.id END) AS likes_count,
    COUNT(CASE WHEN e.dislikes = 1 THEN e.id END) AS dislikes_count
	FROM posts p
	LEFT JOIN category c ON p.id = c.post_id
	LEFT JOIN emotions e ON p.id = e.post_id
    GROUP BY p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, c.name
	ORDER BY p.created_at DESC
	`)
	if err != nil {
		r.log.Printf("error while to query Get All Posts: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		var categ string
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &post.CreatedAt, &post.UpdatedAt, &categ, &post.Likes, &post.Dislikes)
		if err != nil {
			r.log.Printf("error while to scan Get All Posts: %s\n", err.Error())
			return nil, err
		}
		uniqe := false
		post.Category = append(post.Category, categ)

		for i := 0; i < len(posts); i++ {
			if posts[i].Id == post.Id {
				posts[i].Category = append(posts[i].Category, categ)
				uniqe = false
				break
			} else {
				uniqe = true
			}
		}
		if uniqe || len(posts) == 0 {
			posts = append(posts, post)
		}
	}
	if err = rows.Err(); err != nil {
		r.log.Printf("error while to rows.Err() Get All Posts: %s\n", err.Error())
		return nil, err
	}

	return posts, nil
}
