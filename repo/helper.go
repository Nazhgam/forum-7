package repo

import (
	"database/sql"
)

func createTables(db *sql.DB) error {
	// Create User table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE,
			password TEXT,
			email TEXT UNIQUE,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	// Create Post table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			title TEXT,
			content TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}

	// Create Category table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS category (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER ,
			name TEXT,
			FOREIGN KEY (post_id) REFERENCES posts(id)
		)
	`)
	if err != nil {
		return err
	}

	// Create Comment table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			post_id INTEGER,
			content TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (post_id) REFERENCES posts(id)
		)
	`)
	if err != nil {
		return err
	}

	// Create Profile table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS profiles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			name TEXT,
			bio TEXT,
			image_url TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}

	// Create emotions table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS emotions (
		id INTEGER PRIMARY KEY,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL,
		likes BOOLEAN NOT NULL,
		dislikes BOOLEAN NOT NULL,
		FOREIGN KEY (post_id) REFERENCES posts(id)
		FOREIGN KEY (user_id) REFERENCES users(id)
	  )
`)
	if err != nil {
		return err
	}

	return nil
}
