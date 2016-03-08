package newsletters

import (
	"fmt"
	"munch/sources"
	"munch/stories"
)

type Newsletter struct {
	id                  int
	Title               string
	PubDate             int64
	sources             []*sources.SourceManipulator
	sources_lastchecked []int64
}

func New(id int, title string, pubdate int64) *Newsletter {
	ret := new(Newsletter)

	ret.id = id
	ret.Title = title
	ret.PubDate = pubdate

	return ret
}

// Need to call this afterwards
func (newsletter *Newsletter) AddSource(src sources.SourceManipulator, lastchecked int64) {
	newsletter.sources = append(newsletter.sources, &src)
	newsletter.sources_lastchecked = append(newsletter.sources_lastchecked, lastchecked)
}

func (newsletter *Newsletter) PrintNewsletter() {
	fmt.Printf("%s\n========================\n", newsletter.Title)
	fmt.Printf("%d sources:\n", len(newsletter.sources))

	if len(newsletter.sources) > 0 {
		for _, src := range newsletter.sources {
			fmt.Printf("    %s\n", (*src).GetName())
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
	ret := make([]*stories.Story, 0, 5)
	for j, src := range newsletter.sources {
		var latest_pubdate int64 = newsletter.sources_lastchecked[j]
		for _, story := range all_stories {
			if story.GetPubDate() > newsletter.sources_lastchecked[j] &&
				story.GetSourceId() == (*src).GetId() {
				ret = append(ret, story)

				if latest_pubdate < story.GetPubDate() {
					latest_pubdate = story.GetPubDate()
				}
			}
		}
		newsletter.sources_lastchecked[j] = latest_pubdate
	}
	return ret
}

// Updates the "pubdate" field
func (newsletter *Newsletter) MarkPublished() {
	var max_pubdate int64
	max_pubdate = -1
	for _, dt := range newsletter.sources_lastchecked {
		if dt > max_pubdate {
			max_pubdate = dt
		}
	}
	newsletter.PubDate = max_pubdate
}

func (newsletter *Newsletter) GetId() int {
	return newsletter.id
}

func (newsletter *Newsletter) GetTitle() string {
	return newsletter.Title
}

func (newsletter *Newsletter) GetPubDate() int64 {
	return newsletter.PubDate
}

func (newsletter *Newsletter) GetSources() []*sources.SourceManipulator {
	return newsletter.sources
}

func (newsletter *Newsletter) GetSourcesLastChecked() []int64 {
	return newsletter.sources_lastchecked
}
