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
	stories := db.FetchStoriesSince(-1)
	interesting_stories := nl.GetInterestingStories(stories)
	fmt.Printf("Story count = %d\n", len(interesting_stories))
	db.Close()
}
