package database

import (
	"fmt"
	"munch/newsletters"
)

////////////////////////////////////////////////////////////
// Newsletter stuff
func (db *Db) FetchNewsletter(id int) *newsletters.Newsletter {
	output, err := db.connection.Query("SELECT id, title, pubdate FROM broccoli_newsletters WHERE id=?", id)
	if err != nil {
		fmt.Printf("There was a problem with fetching the newsletter: %d\n%s\n", id, err)
	}

	var id_fetched int
	var title string
	var pubdate int64
	output.Scan(&id_fetched, &title, &pubdate)
	return newsletters.NewNewsletter(id_fetched, title, pubdate)
}
