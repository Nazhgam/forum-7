package helper

const (
	SearchPostByCategAndTitle = `SELECT * FROM posts 
	WHERE category LIKE '$1' OR title LIKE '$2'
	ORDER BY COALESCE(category, ''), COALESCE(title, '')`
)
