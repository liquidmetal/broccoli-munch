package sources

import (
	"github.com/ChimeraCoder/anaconda"
	"munch/config"
	"munch/stories"
	"fmt"
	"time"
	"net/url"
	"strconv"
)

type SourceTwitter struct {
	Source
	url string
	cfg *config.Config
}

// Initialize an client library for a given user.
// This only needs to be done *once* per user
func (src *SourceTwitter) InitializeClient() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(src.cfg.GetTwitterConsumerKey())
	anaconda.SetConsumerSecret(src.cfg.GetTwitterConsumerSecret())
	api := anaconda.NewTwitterApi(src.cfg.GetTwitterAccessToken(), src.cfg.GetTwitterAccessSecret())
	fmt.Println(*api.Credentials)
	return api
}


func NewSourceTwitter(id int, name string, url string, lastcrawled int64, cfg *config.Config) *SourceTwitter {
	ret := new(SourceTwitter)
	ret.id = id
	ret.name = name
	ret.url = url
	ret.lastCrawled = lastcrawled
	ret.cfg = cfg

	return ret
}
func (src *SourceTwitter) GetSearch() {

	api := src.InitializeClient()
	search_result, err := api.GetSearch("golang", nil)
	if err != nil {
		panic(err)
	}
	for _, tweet := range search_result.Statuses {
		fmt.Print(tweet.Text)
	}
}

// Throttling queries can easily be handled in the background, automatically
func (src *SourceTwitter) Throttling() {
	api := src.InitializeClient()
	api.EnableThrottling(10*time.Second, 5)

	// These queries will execute in order
	// with appropriate delays inserted only if necessary
	golangTweets, err := api.GetSearch("golang", nil)
	anacondaTweets, err2 := api.GetSearch("anaconda", nil)

	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err)
	}

	fmt.Println(golangTweets)
	fmt.Println(anacondaTweets)
}

// Fetch a list of all followers without any need for managing cursors
// (Each page is automatically fetched when the previous one is read)
func (src *SourceTwitter) GetFollowersListAll() {
	api := src.InitializeClient()
	pages := api.GetFollowersListAll(nil)
	for page := range pages {
		//Print the current page of followers
		fmt.Println(page.Followers)
	}
}


// func (src *SourceTwitter) fetchRecentTweets() (timeline []anaconda.Tweet, err anaconda.error){
// 	api := src.InitializeClient()
// 	v := url.Values{}
// 	v.Set("count", strconv.FormatInt(src.cfg.GetTwitterMaxResultCount(), 10))
// 	return api.GetHomeTimeline(v)
// 	//return result, nil
// }


func (src *SourceTwitter) FetchNewData() []tempArticle {
	api := src.InitializeClient()
	v := url.Values{}
	v.Set("count", strconv.FormatInt(src.cfg.GetTwitterMaxResultCount(), 10))
	items, err := api.GetHomeTimeline(v)
	if err != nil {
		return nil
	}

	var latestCrawl int64 = src.lastCrawled
	var ret []tempArticle
	for _, plitem := range items {
		t := new(tempArticle)

		t.url = fmt.Sprintf("http://twitter.com/%s/status/%s", plitem.User, plitem.IdStr)
		t.content = plitem.Text
		t.description = plitem.User.Description
		t.title = plitem.User.Name
		t.keywords = ""
		t.image = plitem.User.ProfileImageUrlHttps

		pd, err := parse_time(plitem.CreatedAt)
		if err != nil {
			fmt.Printf("There was an error parsing the timestamp: %s\n")
		}
		t.pubdate = pd

		if pd > src.lastCrawled {
			ret = append(ret, *t)
			if pd > latestCrawl {
				latestCrawl = pd
			}
		}
	}

	src.lastCrawled = latestCrawl

	fmt.Printf("Returning %d item\n", len(ret))
	return ret

}

func (src *SourceTwitter) GenerateStories(articles []tempArticle) []stories.Story {
	var ret []stories.Story

	for _, a := range articles {
		story := stories.New(a.title, a.content, a.url, a.pubdate, src.id)
		ret = append(ret, *story)
	}

	return ret
}

func (src *SourceTwitter) GetType() int {
	return TypeTwitter
}

func (src *SourceTwitter) GetUrl() string {
	return src.url
}

func (src *SourceTwitter) GetLastCrawled() int64 {
	return src.lastCrawled
}
