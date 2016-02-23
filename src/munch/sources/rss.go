// The sources package
package sources

import (
	"fmt"
	"github.com/advancedlogic/GoOse"
	"io/ioutil"
	"munch/stories"
	"net/http"
	"time"
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

func NewSourceRss(name string, url string, lastcrawled int) *SourceRss {
	ret := new(SourceRss)
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
		parsed_time, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			fmt.Printf("There was an error when trying to figure out the date\n")
			fmt.Printf("%s\n", err)
			continue
		}
		t.pubdate = parsed_time.Unix()
		ret = append(ret, *t)
	}
	return ret
}

func (rss *SourceRss) constructStories(articles []tempArticle) []stories.Story {
	var ret []stories.Story

	for i, _ := range articles {
		article := articles[i]
		story := stories.NewStory(article.title, article.content, article.url, article.pubdate)
		ret = append(ret, *story)
	}
	return ret
}

func (rss *SourceRss) FetchNewData() []tempArticle {
	rss.fetchFeed()
	rss.removeOldItems()
	articles := rss.fetchFullArticles()
	return articles
}

func (rss *SourceRss) GenerateStories(articles []tempArticle) []stories.Story {
	stories := rss.constructStories(articles)
	return stories
}
