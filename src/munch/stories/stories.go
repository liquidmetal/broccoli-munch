package stories

import "fmt"

type Story struct {
	Title   string
	Content string
	Link    string
	PubDate int64
}

func NewStory(title string, content string, link string, pubdate int64) *Story {
	story := new(Story)
	story.Title = title
	story.Content = content
	story.Link = link
	story.PubDate = pubdate

	return story
}

func (story *Story) PrintStory() {
	fmt.Printf("%s\n=========================\n%s\n%s\n", story.Title, story.Content, story.Link)
}
