package stories

import "fmt"

type Story struct {
	Title   string
	Summary string
	Link    string
	PubDate int64
}

func NewStory(title string, summary string, link string, pubdate int64) *Story {
	story := new(Story)
	story.Title = title
	story.Summary = summary
	story.Link = link
	story.PubDate = pubdate

	return story
}

func (story *Story) PrintStory() {
	fmt.Printf("%s\n=========================\n%s\n%s\n", story.Title, story.Summary, story.Link)
}
