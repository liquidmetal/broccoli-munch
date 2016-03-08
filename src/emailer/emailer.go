package main

import (
	"fmt"
	"github.com/mailgun/mailgun-go"
	"munch/database"
	"munch/producer"
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
	//nl.MarkPublished()
	db.PersistNewsletter(nl)

	html := producer.Produce(nl, interesting_stories)
	fmt.Printf("Email HTML:\n%s\n", html)

	db.Close()
}
