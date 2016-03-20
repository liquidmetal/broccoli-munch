package sources

import (
	"encoding/gob"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
	"hash/fnv"
	"munch/config"
	"munch/stories"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type SourceYoutube struct {
	Source
	url string
	cfg *config.Config
}

func NewSourceYoutube(id int, name string, url string, lastcrawled int64, cfg *config.Config) *SourceYoutube {
	ret := new(SourceYoutube)
	ret.id = id
	ret.name = name
	ret.url = url
	ret.lastCrawled = lastcrawled
	ret.cfg = cfg

	return ret
}

func (src *SourceYoutube) fetchRecentUploads() ([]*youtube.PlaylistItem, error) {
	// OAuth configuration
	oauthconfig := &oauth2.Config{
		ClientID:     src.cfg.GetYoutubeClientId(),
		ClientSecret: src.cfg.GetYoutubeClientSecret(),
		Endpoint:     google.Endpoint,
		Scopes:       []string{youtube.YoutubeScope},
	}
	ctx := context.Background()

	// Try connecting to Youtube
	client := newOAuthClient(ctx, oauthconfig, src.cfg)
	service, err := youtube.New(client)
	cs := youtube.NewChannelsService(service)
	call_channel := cs.List("contentDetails").MaxResults(1).Id(src.url)
	resp_channel, err := call_channel.Do()

	if err != nil {
		// The token expired - delete it and try again
		fmt.Printf("%s\n", err)
		fmt.Printf("The token has probably expired. Trying to fetch a new token\n")

		removeCacheFile(oauthconfig)
		client = newOAuthClient(ctx, oauthconfig, src.cfg)

		service, err = youtube.New(client)
		cs := youtube.NewChannelsService(service)
		call_channel := cs.List("contentDetails").MaxResults(1).Id(src.url)
		resp_channel, err = call_channel.Do()
	}

	if err != nil {
		fmt.Printf("%s\n", err)
		return nil, errors.New("There was an error connecting to Youtube")
	}

	if len(resp_channel.Items) == 0 {
		fmt.Printf("The REST call was successful there was an error fetching the channel ID: %s\n", src.url)
		return nil, errors.New("There was an error connecting to Youtube")
	}

	uploads_playlist := resp_channel.Items[0].ContentDetails.RelatedPlaylists.Uploads

	pl := youtube.NewPlaylistItemsService(service)
	call_playlist := pl.List("snippet").MaxResults(src.cfg.GetYoutubeMaxResultCount()).PlaylistId(uploads_playlist)
	resp_playlist, err := call_playlist.Do()

	return resp_playlist.Items, nil
}

func (src *SourceYoutube) FetchNewData() []tempArticle {
	items, err := src.fetchRecentUploads()
	if err != nil {
		return nil
	}

	var latestCrawl int64 = src.lastCrawled
	var ret []tempArticle
	for _, plitem := range items {
		t := new(tempArticle)

		t.url = fmt.Sprintf("http://youtube.com/watch?v=%s", plitem.Snippet.ResourceId.VideoId)
		t.content = plitem.Snippet.Description
		t.description = plitem.Snippet.Description
		t.title = plitem.Snippet.Title
		t.keywords = ""
		t.image = plitem.Snippet.Thumbnails.High.Url

		pd, err := parse_time(plitem.Snippet.PublishedAt)
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

func (src *SourceYoutube) GetType() int {
	return TypeYoutube
}

func (src *SourceYoutube) GetUrl() string {
	return src.url
}

func (src *SourceYoutube) GenerateStories(articles []tempArticle) []stories.Story {
	var ret []stories.Story

	for _, a := range articles {
		story := stories.New(a.title, a.content, a.url, a.pubdate, src.id)
		ret = append(ret, *story)
	}

	return ret
}

func (src *SourceYoutube) GetLastCrawled() int64 {
	return src.lastCrawled
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := new(oauth2.Token)
	err = gob.NewDecoder(f).Decode(t)
	return t, err
}

func newOAuthClient(ctx context.Context, config *oauth2.Config, cfg *config.Config) *http.Client {
	cacheFile := tokenCacheFile(config)
	token, err := tokenFromFile(cacheFile)
	if err != nil {
		token = tokenFromWeb(ctx, config, cfg)
		token.RefreshToken = cfg.GetYoutubeRefreshToken()
		saveToken(cacheFile, token)
	} else {
		fmt.Printf("Using cached token from %q\n", cacheFile)
	}

	return config.Client(ctx, token)
}

func tokenCacheFile(config *oauth2.Config) string {
	hash := fnv.New32a()
	hash.Write([]byte(config.ClientID))
	hash.Write([]byte(config.ClientSecret))
	hash.Write([]byte(strings.Join(config.Scopes, " ")))
	fn := fmt.Sprintf("go-api-demo-tok%v", hash.Sum32())
	return filepath.Join(osUserCacheDir(), url.QueryEscape(fn))
}

func tokenFromWeb(ctx context.Context, config *oauth2.Config, cfg *config.Config) *oauth2.Token {
	ch := make(chan string)
	randState := fmt.Sprintf("st%d", time.Now().UnixNano())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GetYoutubeOAuthPort()))
	if err != nil {
		panic(err)
	}
	// make our own handler that puts all requests in a wait group.
	h := http.NewServeMux()

	// add a close handlefunc to that handler
	h.HandleFunc("/youtube_callback/", func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/favicon.ico" {
			http.Error(rw, "", 404)
			return
		}
		if req.FormValue("state") != randState {
			fmt.Printf("State doesn't match: req = %#v", req)
			http.Error(rw, "", 500)
			return
		}
		if code := req.FormValue("code"); code != "" {
			fmt.Fprintf(rw, "<h1>Success</h1>Authorized.")
			rw.(http.Flusher).Flush()
			ch <- code
			lis.Close()
			return
		}
		fmt.Printf("no code")
		http.Error(rw, "", 500)
		lis.Close()
	})

	//listen and serve until listner is closed
	go http.Serve(lis, h)

	// TODO fix this hard-coded localhost thingy
	config.RedirectURL = fmt.Sprintf("http://127.0.0.1:%d/youtube_callback/", cfg.GetYoutubeOAuthPort())
	authURL := config.AuthCodeURL(randState)
	go openURL(authURL)
	fmt.Printf("Authorize this app at: %s", authURL)
	code := <-ch

	token, err := config.Exchange(ctx, code)
	if err != nil {
		fmt.Printf("Token exchange error: %v", err)
	}
	return token
}

func openURL(url string) {
	try := []string{"xdg-open", "google-chrome", "open"}
	for _, bin := range try {
		err := exec.Command(bin, url).Run()
		if err == nil {
			return
		}
	}
	fmt.Printf("Error opening URL in browser.")
}

func saveToken(file string, token *oauth2.Token) {
	f, err := os.Create(file)
	if err != nil {
		fmt.Printf("Warning: failed to cache oauth token: %v", err)
		return
	}
	defer f.Close()
	fmt.Printf("AccessToken = %s\n", token.AccessToken)
	fmt.Printf("TokenType = %s\n", token.TokenType)
	fmt.Printf("RefreshToken = %s\n", token.RefreshToken)
	fmt.Printf("Expiry = %s\n", token.Expiry)
	gob.NewEncoder(f).Encode(token)
}

func removeCacheFile(cfg *oauth2.Config) error {
	fname := tokenCacheFile(cfg)
	fmt.Printf("Removing: %s\n", fname)
	err := os.Remove(fname)
	return err
}

func osUserCacheDir() string {
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Caches")
	case "linux", "freebsd":
		return filepath.Join(os.Getenv("HOME"), ".cache")
	}
	fmt.Printf("TODO: osUserCacheDir on GOOS %q", runtime.GOOS)
	return "."
}
