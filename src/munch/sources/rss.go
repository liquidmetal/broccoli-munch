// The sources package
package sources

import (
	"fmt"
	"github.com/advancedlogic/GoOse"
	"io/ioutil"
	"net/http"
)

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
	for i, _ := range rss.data.ItemList {
		// TODO remove old items
		fmt.Printf("%s %s\n", rss.data.ItemList[i].Link, rss.data.ItemList[i].PubDate)
	}
}

func (rss *SourceRss) fetchFullArticles() []string {
	ret := make([]string, 10, 10)
	for i, _ := range rss.data.ItemList {
		item := rss.data.ItemList[i]
		//response, err := http.Get(item.Link)
		//if err != nil {
		//		fmt.Printf("There was an error when fetching the article: %s\n", item.Link)
		//}
		//contents, err := ioutil.ReadAll(response.Body)

		g := goose.New()
		article, _ := g.ExtractFromURL(item.Link)
		fmt.Printf("%s: %s\n%s\n=======================================================\n",
			item.Title,
			article.FinalURL,
			string(article.MetaDescription))
	}
	return ret
}

func (rss *SourceRss) FetchNewData() {
	rss.fetchFeed()
	rss.removeOldItems()
	rss.fetchFullArticles()
}

func (rss *SourceRss) GenerateStories() {
}

func (rss *SourceRss) PersistStories() {
}
