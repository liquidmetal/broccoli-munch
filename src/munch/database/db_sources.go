package database

import (
	"fmt"
	"munch/sources"
)

/////////////////////////////////////////////////////////////////
// Functions for sources

func (db *Db) PersistSource(src sources.SourceManipulator) {
	id := src.GetId()
	fmt.Printf("Trying to update: %d\n", id)
	name := src.GetName()
	url := src.GetUrl()
	src_type := src.GetType()
	lc := src.GetLastCrawled()

	stmt, err := db.connection.Prepare("UPDATE broccoli_sources SET name=?, url=?, type=?, lastcrawled=? WHERE id=?")
	if err != nil {
		fmt.Printf("There was an error when trying to create statement for persisting a source\n%s\n", err)
		return
	}
	res, err := stmt.Exec(name, url, src_type, lc, id)
	if err != nil {
		fmt.Printf("There was an error trying to update the source")
		return
	}

	affect, err := res.RowsAffected()
	if affect != 1 {
		fmt.Printf("An unexpected number of rows were affected: %d vs 1\n", affect)
	}

	fmt.Printf("The source was persisted as expected\n")
}

func (db *Db) FetchSource(id int) (src sources.SourceManipulator) {
	output, err := db.connection.Query("SELECT id, name, url, type, lastcrawled FROM broccoli_sources WHERE id = ?;", id)
	if err != nil {
		fmt.Printf("There was an error fetching sources\n%s\n", err)
	}

	defer output.Close()

	var id_fetched int
	var name string
	var url string
	var source_type int
	var last_crawled int64
	output.Next()
	output.Scan(&id_fetched, &name, &url, &source_type, &last_crawled)

	if id_fetched != id {
		fmt.Printf("There was an error with the ID that was fetched (%d vs %d)\n", id_fetched, id)
	}
	fmt.Printf("The source URL is %s\n", url)

	if source_type == sources.TypeRss {
		return sources.NewSourceRss(id, name, url, last_crawled)
	} else if source_type == sources.TypeYoutube {
		return sources.NewSourceYoutube(id, name, url, last_crawled, db.cfg)
	} else if source_type == sources.TypeTwitter {
		return sources.NewSourceTwitter(id, name, url, last_crawled, db.cfg)
	} else {
		fmt.Printf("Unable to match source type\n")
	}

	return nil
}

func (db *Db) FetchAllSourceIds() ([]int64, error) {
	output, err := db.connection.Query("SELECT id FROM broccoli_sources")
	if err != nil {
		fmt.Printf("There was an error fetching a list of all sources\n%s\n", err)
		return nil, fmt.Errorf("There was an error")
	}

	var id int64
	var ret []int64
	for output.Next() {
		output.Scan(&id)
		ret = append(ret, id)
	}

	return ret, nil
}
