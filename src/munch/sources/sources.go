// The sources package
package sources

import (
	"errors"
	"fmt"
	"munch/stories"
	"time"
)

const (
	TypeRss     = iota
	TypeBlog    = iota
	TypeYoutube = iota
	TypeGithub  = iota
)

type SourceManipulator interface {
	FetchNewData() []tempArticle
	GenerateStories(articles []tempArticle) []stories.Story
	GetId() int
	GetName() string
	GetType() int
	GetUrl() string
	GetLastCrawled() int64
}

type Source struct {
	id          int
	name        string
	lastCrawled int64
}

type SourceRss struct {
	Source
	url  string
	data Rss2
}

type tempArticle struct {
	url         string
	content     string
	description string
	title       string
	keywords    string
	image       string
	pubdate     int64
}

// Utility function for parsing any kind of time
// TODO properly arrange these based on the kind of traffic
func parse_time(mtime string) (int64, error) {
	temp, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", mtime)
	if err == nil {
		return temp.Unix(), nil
	}

	temp, err = time.Parse("Mon, 2 Jan 2006 15:04:05 MST", mtime)
	if err == nil {
		return temp.Unix(), nil
	}

	temp, err = time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", mtime)
	if err == nil {
		return temp.Unix(), nil
	}

	temp, err = time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", mtime)
	if err == nil {
		return temp.Unix(), nil
	}

	temp, err = time.Parse("2006-01-02T15:04:05.000Z", mtime)
	if err == nil {
		return temp.Unix(), nil
	}

	fmt.Printf("There was an error parsing: %s\n", mtime)
	return -1, errors.New("There was an error parsing the timestamp")
}
