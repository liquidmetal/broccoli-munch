package main

import (
	"fmt"
	"github.com/mailgun/mailgun-go"
	"munch/config"
	"munch/database"
	"munch/producer"
)

func main() {
	fmt.Printf("This is the emailer\n")

	cfg := config.New()
	db := database.New(cfg)
	nl := db.FetchNewsletter(1)
	db.FetchNewsletterSources(nl)
	nl.PrintNewsletter()
	stories := db.FetchStoriesSince(nl.GetPubDate())
	fmt.Printf("Original story count = %d\n", len(stories))
	interesting_stories := nl.GetInterestingStories(stories)
	fmt.Printf("Story count = %d\n", len(interesting_stories))
	//nl.MarkPublished()
	db.PersistNewsletter(nl)

	html := producer.Produce(nl, interesting_stories)

	// Uncomment these lines to enable sending emails
	//cfg := config.New()
	//send_email("sinha.utkarsh1990@gmail.com", nl.Title, html, cfg)
	fmt.Printf("%s\n", html)

	db.Close()
}

func send_email(email string, title string, html string, config *config.Config) {
	gun := mailgun.NewMailgun("munch.utkarshsinha.com",
		config.GetMailPrivateKey(),
		config.GetMailPublicKey())
	msg := mailgun.NewMessage("Broccoli Munch <noreplies@munch.utkarshsinha.com>",
		title,
		html,
		"sinha.utkarsh1990@gmail.com")
	msg.SetHtml(html)
	response, id, err := gun.Send(msg)

	fmt.Printf("%s\n", response)
	fmt.Printf("%s\n", id)
	fmt.Printf("%s\n", err)
}
