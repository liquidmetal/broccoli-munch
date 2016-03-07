// The sources package
package sources

import "munch/stories"

const (
	TypeRss    = iota
	TypeBlog   = iota
	TypeVideo  = iota
	TypeGithub = iota
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
