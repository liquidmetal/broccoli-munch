package main

import (
	"fmt"
	"munch/config"
	"munch/database"
	"munch/messaging"
	"munch/stories"
)

// Worker process for fetching data from a given source

func main() {
	// Step 1: Establish connection to the message queue
	// Step 2: Wait for a task to show up in the queue
	// Step 3: Fetch the data for the given task, process it, generate stories
	// Step 4: Put the data into a DB somewhere
	// Step 5: Repeat

	cfg := config.New()
	broker := messaging.New(cfg)

	for true {
		fmt.Printf("Waiting for crawl request...\n")
		sourceid := broker.DequeueCrawl()
		db := database.New(cfg)
		source := db.FetchSource(sourceid)
		fmt.Printf("Fetching %s...\n", source.GetName())

		articles := source.FetchNewData()

		if articles != nil {
			s := source.GenerateStories(articles)
			persistStories(s, db)
		} else {
			fmt.Println("Yo something happened here\n")
		}
		db.PersistSource(source)
		db.Close()
	}
}

func persistStories(stories []stories.Story, db *database.Db) {
	count := 0
	for _, story := range stories {
		// If the story doesn't already exist, persist it
		if !db.StoryExists(&story) {
			db.PersistStory(&story)
			count += 1
		}
	}
	fmt.Printf("Saved %d new stories\n", count)
}
