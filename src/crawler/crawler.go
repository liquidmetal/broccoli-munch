package main

import (
	"fmt"
	"munch/database"
	"munch/stories"
	"os"
	"strconv"
)

// Worker process for fetching data from a given source

func main() {
	// Step 1: Establish connection to the message queue
	// Step 2: Wait for a task to show up in the queue
	// Step 3: Fetch the data for the given task, process it, generate stories
	// Step 4: Put the data into a DB somewhere
	// Step 5: Repeat

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <id>\n", os.Args[0])
		return
	}

	sourceid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("There was an error converting source ID to integer (%s)\n", sourceid)
		return
	}

	db := database.NewDB()
	fmt.Printf("Fetching techcrunch...\n")
	source := db.FetchSource(sourceid)

	//f := sources.NewSourceRss("Techcrunch", "http://feeds.feedburner.com/TechCrunch/?fmt=xml", 0)
	//f := sources.NewSourceRss("Utkarsh Sinha", "http://utkarshsinha.com/index.xml", 0)
	articles := source.FetchNewData()
	s := source.GenerateStories(articles)
	s[0].PrintStory()

	persistStories(s, db)

	db.PersistSource(source)
	db.Close()
}

func persistStories(stories []stories.Story, db *database.Db) {
	for _, story := range stories {
		// If the story doesn't already exist, persist it
		if !db.StoryExists(&story) {
			db.PersistStory(&story)
		}
	}
}
