// The sources package
package sources

const (
	typeRss    = iota
	typeBlog   = iota
	typeVideo  = iota
	typeGithub = iota
)

type SourceManipulator interface {
	FetchNewData()
	GenerateStories()
}

type Source struct {
	id          int
	name        string
	lastCrawled int
}

type SourceRss struct {
	Source
	url  string
	data Rss2
}
