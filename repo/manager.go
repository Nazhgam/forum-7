package repo

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type IDb interface {
	IUser
	IPost
	IComment
	IProfile
	ICateg
	IEmotion
}

type repo struct {
	db  *sql.DB
	log *log.Logger
}

func New(l *log.Logger) (IDb, error) {
	// Open SQLite3 database
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	err = createTables(db)
	if err != nil {
		log.Printf("error while create table of database %s\n", err.Error())
		return nil, err
	}

	return repo{db: db, log: l}, nil
}
