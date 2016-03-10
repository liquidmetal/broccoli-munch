package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const (
	TABLE_STORIES = "broccoli_stories"
	TABLE_SOURCES = "broccoli_sources"
)

type Db struct {
	connection *sql.DB
}

///////////////////////////////////////////////////////////////
// Database housekeeping
func New(filename string) *Db {
	db := new(Db)
	connection, err := sql.Open("sqlite3", filename)

	if err != nil {
		fmt.Printf("There was an error opening the database\n")
		return nil
	}
	db.connection = connection
	return db
}

func (db *Db) Close() {
	err := db.connection.Close()
	if err != nil {
		fmt.Printf("There was an error closing the database: %s\n", err)
	}
}
