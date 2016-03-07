package newsletters

import (
	"munch/stories"
)

type Newsletter struct {
	id      int
	Title   string
	PubDate int64
}

func NewNewsletter(id int, title string, pubdate int64) *Newsletter {
	ret := new(Newsletter)

	ret.id = id
	ret.Title = title
	ret.PubDate = pubdate

	return ret
}

// Generates an array of interesting stories
func (newsletter *Newsletter) GetInterestingStories() []stories.Story {
	return nil
}
