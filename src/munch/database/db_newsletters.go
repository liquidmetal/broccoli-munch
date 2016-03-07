package database

import (
	"fmt"
	"munch/newsletters"
	"munch/users"
)

////////////////////////////////////////////////////////////
// Newsletter stuff
func (db *Db) FetchNewsletter(id int) *newsletters.Newsletter {
	output, err := db.connection.Query("SELECT id, title, pubdate FROM broccoli_newsletters WHERE id=?", id)
	if err != nil {
		fmt.Printf("There was a problem with fetching the newsletter: %d\n%s\n", id, err)
	}

	output.Next()
	defer output.Close()

	var id_fetched int
	var title string
	var pubdate int64
	output.Scan(&id_fetched, &title, &pubdate)
	return newsletters.New(id_fetched, title, pubdate)
}

// WARNING: Mutation of newsletter here
func (db *Db) FetchNewsletterSources(newsletter *newsletters.Newsletter) {
	output, err := db.connection.Query("SELECT source_id, source_lastchecked FROM broccoli_newsletters_sources WHERE newsletter_id = ?", newsletter.GetId())
	if err != nil {
		fmt.Printf("There was an error fetching all the source IDs\n%s\n", err)
	}

	for output.Next() {
		var source_id int
		var source_lastchecked int64
		output.Scan(&source_id, source_lastchecked)
		source := db.FetchSource(source_id)
		newsletter.AddSource(source, source_lastchecked)
	}
}

func (db *Db) FetchUserNewsletters(user users.User) []newsletters.Newsletter {
	// TODO this can probably be converted into a join
	output, err := db.connection.Query("SELECT newsletter_id FROM broccoli_users_newsletters WHERE user_id=?", user.GetId())
	defer output.Close()

	if err != nil {
		fmt.Printf("There was an error trying to fetch the newsletters of a person:\n%s\n", err)
	}

	return nil
}
