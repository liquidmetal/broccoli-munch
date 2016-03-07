package database

import (
	"fmt"
	"munch/stories"
)

///////////////////////////////////////////////////////////////
// Functions for stories

func (db *Db) StoryExists(story *stories.Story) (exists bool) {
	title := story.Title
	url := story.Link
	pubdate := story.PubDate
	source_id := story.SourceId
	res, err := db.connection.Query("SELECT COUNT(id) FROM broccoli_stories WHERE title=? AND url=? AND pubdate=? AND source_id=?", title, url, pubdate, source_id)
	if err != nil {
		fmt.Printf("There was an error in fetching existing stories")
	}
	defer res.Close()

	var count int
	res.Next()
	res.Scan(&count)
	if count == 0 {
		return false
	}
	return true
}

func (db *Db) PersistStory(story *stories.Story) {
	stmt, err := db.connection.Prepare("INSERT INTO broccoli_stories(url, title, content, source_id, pubdate) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("There was an error when constructing the statement\n%s", err)
		return
	}
	_, err = stmt.Exec(story.Link, story.Title, story.Content, story.SourceId, story.PubDate)
	if err != nil {
		fmt.Printf("There was an error when trying to persist the story\n%s", err)
		return
	}
}

func (db *Db) FetchStoriesSince(timestamp int64) []*stories.Story {
	result, err := db.connection.Query("SELECT id, url, title, content, pubdate, source_id FROM broccoli_stories WHERE pubdate >= ?", timestamp)

	if err != nil {
		fmt.Printf("There was an error when fetching stories since: %s\n", err)
	}

	ret := make([]*stories.Story, 0, 10)
	for result.Next() {
		var id int
		var url string
		var title string
		var content string
		var pubdate int64
		var source_id int
		result.Scan(&id, &url, &title, &content, &pubdate, &source_id)
		ret = append(ret, stories.New(title, content, url, pubdate, source_id))
	}
	return ret
}
