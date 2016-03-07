package database

import (
	"fmt"
	"munch/users"
)

func (db *Db) FetchUser(id int) *users.User {
	result, err := db.connection.Query("SELECT (id, name, email) FROM broccoli_users WHERE id=?", id)
	if err != nil {
		fmt.Printf("There was an error when trying to fetch the user:\n%s\n", err)
	}

	var id_fetched int
	var name string
	var email string
	result.Scan(&id_fetched, &name, &email)

	return users.NewUser(id_fetched, name, email)
}
