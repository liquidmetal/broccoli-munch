package newsletters

import (
	"fmt"
	"munch/sources"
	"munch/stories"
)

type Newsletter struct {
	id                 int
	Title              string
	PubDate            int64
	sources            []sources.SourceManipulator
	source_lastchecked []int64
}

func New(id int, title string, pubdate int64) *Newsletter {
	ret := new(Newsletter)

	ret.id = id
	ret.Title = title
	ret.PubDate = pubdate

	return ret
}

func (newsletter *Newsletter) AddSource(src sources.SourceManipulator, lastchecked int64) {
	newsletter.sources = append(newsletter.sources, src)
	newsletter.source_lastchecked = append(newsletter.source_lastchecked, lastchecked)
}

func (newsletter *Newsletter) PrintNewsletter() {
	fmt.Printf("%s\n========================\n", newsletter.Title)
	fmt.Printf("%d sources:\n", len(newsletter.sources))

	if len(newsletter.sources) > 0 {
		for _, src := range newsletter.sources {
			fmt.Printf("    %s\n", src.GetName())
		}
	}
}

// Given a bunch of recent stories (fetched from the db) generates a list
// of stories that will actually make it to the final newsletter
func (newsletter *Newsletter) GetInterestingStories(all_stories []*stories.Story) []*stories.Story {
	// No sources setup for this? Return nothing
	if len(newsletter.sources) == 0 {
		return nil
	}

	// Fetchall stories after the published source (for every source)
	ret := make([]*stories.Story, 0, 10)
	for _, story := range all_stories {
		for _, src := range newsletter.sources {
			if story.GetSourceId() == src.GetId() {
				ret = append(ret, story)
				// TODO update the source_lastchecked field
			}
		}
	}
	return ret
}

func (newsletter *Newsletter) GetId() int {
	return newsletter.id
}
