package stories

import (
	"fmt"
	"github.com/JesusIslam/tldr"
)

type Story struct {
	Title    string
	Content  string
	Link     string
	PubDate  int64
	SourceId int
}

func NewStory(title string, content string, link string, pubdate int64, src int) *Story {
	story := new(Story)
	story.Title = title
	story.Content = content
	story.Link = link
	story.PubDate = pubdate
	story.SourceId = src

	return story
}

func (story *Story) PrintStory() {
	fmt.Printf("%s\n=========================\n%s\n%s\n", story.Title, story.Summarize(1), story.Link)
}

func (story *Story) Summarize(sentences int) string {
	tldr := tldr.New()
	result, err := tldr.Summarize(story.Content, sentences)
	if err != nil {
		fmt.Printf("There was an error summarizing text:\n%s\n", err)
	}

	return result
}

}
