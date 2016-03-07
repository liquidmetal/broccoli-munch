package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"munch/sources"
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
	stmt, err := db.connection.Prepare("INSERT INTO broccoli_stories(url, title, content, pubdate) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("There was an error when constructing the statement\n%s", err)
		return
	}
	_, err = stmt.Exec(story.Link, story.Title, story.Content, story.PubDate)
	if err != nil {
		fmt.Printf("There was an error when trying to persist the story\n%s", err)
		return
	}
}

func (db *Db) PersistSource(src sources.SourceManipulator) {
	id := src.GetId()
	fmt.Printf("Trying to update: %d\n", id)
	name := src.GetName()
	url := src.GetUrl()
	src_type := src.GetType()
	lc := src.GetLastCrawled()

	stmt, err := db.connection.Prepare("UPDATE broccoli_sources SET name=?, url=?, type=?, lastcrawled=? WHERE id=?")
	if err != nil {
		fmt.Printf("There was an error when trying to create statement for persisting a source\n%s\n", err)
		return
	}
	res, err := stmt.Exec(name, url, src_type, lc, id)
	if err != nil {
		fmt.Printf("There was an error trying to update the source")
		return
	}

	affect, err := res.RowsAffected()
	if affect != 1 {
		fmt.Printf("An unexpected number of rows were affected: %d vs 1\n", affect)
	}

	fmt.Printf("The source was persisted as expected\n")
}

func (db *Db) StoryExists(story *stories.Story) (exists bool) {
	return false
}

func (db *Db) FetchSource(id int) (src sources.SourceManipulator) {
	output, err := db.connection.Query("SELECT id, name, url, type, lastcrawled FROM broccoli_sources WHERE id = 1;")
	if err != nil {
		fmt.Printf("There was an error fetching sources\n%s\n", err)
	}

	defer output.Close()

	var id_fetched int
	var name string
	var url string
	var source_type int
	var last_crawled int64
	output.Next()
	output.Scan(&id_fetched, &name, &url, &source_type, &last_crawled)

	if id_fetched != id {
		fmt.Printf("There was an error with the ID that was fetched (%d vs %d)\n", id_fetched, id)
	}
	fmt.Printf("The source URL is %s\n", url)

	if source_type == sources.TypeRss {
		return sources.NewSourceRss(id, name, url, last_crawled)
	} else {
		fmt.Printf("Unable to match source type")
	}

	return nil
}

func (db *Db) Close() {
	err := db.connection.Close()
	if err != nil {
		fmt.Printf("There was an error closing the database: %s\n", err)
	}
}
