// The sources package
package sources

import (
	"fmt"
	"github.com/advancedlogic/GoOse"
	"io/ioutil"
	"munch/stories"
	"net/http"
)

type tempArticle struct {
	url         string
	content     string
	description string
	title       string
	keywords    string
	image       string
	pubdate     int64
}

func NewSourceRss(id int, name string, url string, lastcrawled int64) *SourceRss {
	ret := new(SourceRss)
	ret.id = id
	ret.name = name
	ret.url = FixRssUrl(url)
	ret.lastCrawled = lastcrawled

	return ret
}

func FixRssUrl(url string) string {
	// TODO If this is a feedburner link, ensure there is a ?fmt=xml at the end
	// TODO Ensure that the URL starts with http://<domain>.com
	return url
}

func (rss *SourceRss) fetchFeed() {
	response, err := http.Get(rss.url)
	if err != nil {
		fmt.Printf("There was an error when fetching %s\n", rss.url)
		fmt.Printf("%s\n", err)
		return
	}

	mtime := response.Header.Get("Last-Modified")
	parsed_time, err := parse_time(mtime)
	rss.lastCrawled = parsed_time

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("There was an error when reading the response of %s\n", rss.url)
		fmt.Printf("%s\n", err)
		return
	}

	parsed, failed := parseFeedContent(contents)
	if failed != true {
		fmt.Printf("There was an error parsing the contents\n")
	}

	rss.data = parsed
}

// Remove old items from the list of items. This is based on the timestamp
// of the last crawl
func (rss *SourceRss) removeOldItems() {
	fmt.Printf("%s\n", rss.data.Title)
	/*for i, _ := range rss.data.ItemList {
		// TODO remove old items
	}*/
}

func (rss *SourceRss) fetchFullArticles() []tempArticle {
	var ret []tempArticle
	g := goose.New()
	for i, _ := range rss.data.ItemList {
		item := rss.data.ItemList[i]
		article, err := g.ExtractFromURL(item.Link)
		if err != nil {
			fmt.Printf("There was an error fetching the article for: %s\n", item.Link)
			continue
		}

		t := new(tempArticle)
		t.url = article.FinalURL
		t.content = article.CleanedText
		t.description = article.MetaDescription
		t.keywords = article.MetaKeywords
		t.image = article.TopImage
		t.title = item.Title
		parsed_time, err := parse_time(item.PubDate)
		if err != nil {
			fmt.Printf("There was an error when trying to figure out the date\n")
			fmt.Printf("%s\n", err)
			continue
		}
		t.pubdate = parsed_time
		ret = append(ret, *t)
	}
	return ret
}

func (rss *SourceRss) constructStories(articles []tempArticle) []stories.Story {
	var ret []stories.Story

	for i, _ := range articles {
		article := articles[i]
		story := stories.New(article.title, article.content, article.url, article.pubdate, rss.id)
		ret = append(ret, *story)
	}
	return ret
}

func (rss *SourceRss) hasNewStories() (has bool) {
	response, err := http.Head(rss.url)
	if err != nil {
		fmt.Printf("There was an error running the HEAD command on url:\n%s", rss.url)
		fmt.Printf("Assuming new stories were available\n")
		return true
	}

	// Try and figure out the mtime
	mtime := response.Header.Get("Last-Modified")
	var parsed_time int64
	if len(mtime) > 0 {
		parsed_time, err = parse_time(mtime)
		if err != nil {
			fmt.Printf("There was an error parsing the timestamp: %s\n", mtime)
		}
		fmt.Printf("Last crawled = %d\n", rss.lastCrawled)
		fmt.Printf("Updated = %d (%s)\n", parsed_time, mtime)
	} else {
		fmt.Printf("There was omething wrong with Last-Modified")
	}

	// If there are changes, yes
	if rss.lastCrawled < parsed_time {
		return true
	}
	return false
}

func (rss *SourceRss) FetchNewData() []tempArticle {
	has_new := rss.hasNewStories()

	if has_new {
		rss.fetchFeed()
		rss.removeOldItems()
		articles := rss.fetchFullArticles()
		return articles
	}

	return nil
}

func (rss *SourceRss) GenerateStories(articles []tempArticle) []stories.Story {
	stories := rss.constructStories(articles)
	return stories
}

func (rss *SourceRss) GetId() int {
	return rss.id
}

func (rss *SourceRss) GetName() string {
	return rss.name
}

func (rss *SourceRss) GetUrl() string {
	return rss.url
}

func (rss *SourceRss) GetType() int {
	return TypeRss
}

func (rss *SourceRss) GetLastCrawled() int64 {
	return rss.lastCrawled
}
