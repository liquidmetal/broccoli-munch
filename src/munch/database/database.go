package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"munch/stories"
)

const (
	TABLE_STORIES = "broccoli_stories"
	TABLE_SOURCES = "broccoli_sources"
)

type Db struct {
	connection *sql.DB
}

func NewDB() *Db {
	db := new(Db)
	connection, err := sql.Open("sqlite3", "./broccoli.db")

	if err != nil {
		fmt.Printf("There was an error opening the database\n")
		return nil
	}
	db.connection = connection
	return db
}

func (db *Db) PersistStory(story *stories.Story) {
}

func (db *Db) Close() {
	db.connection.Close()
}
