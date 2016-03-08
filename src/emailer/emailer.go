package main

import (
	"fmt"
	"munch/database"
)

func main() {
	fmt.Printf("This is the emailer\n")

	db := database.NewDB()
	nl := db.FetchNewsletter(1)
	db.FetchNewsletterSources(nl)
	nl.PrintNewsletter()
	stories := db.FetchStoriesSince(nl.GetPubDate())
	fmt.Printf("Original story count = %d\n", len(stories))
	interesting_stories := nl.GetInterestingStories(stories)
	fmt.Printf("Story count = %d\n", len(interesting_stories))
	for _, s := range interesting_stories {
		s.PrintStory()
	}
	nl.MarkPublished()
	db.PersistNewsletter(nl)
	db.Close()
}
