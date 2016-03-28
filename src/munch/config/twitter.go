package config

import "strconv"
import "fmt"

func (config *Config) parseConfigTwitterValues(t *ConfigReader) {
	// Mail settings
	config.twitter.consumerkey = t.Twitter["consumerkey"]
	config.twitter.consumersecret = t.Twitter["clientsecret"]
	config.twitter.accesstoken = t.Twitter["accesstoken"]
	config.twitter.accesssecret = t.Twitter["accesssecret"]

	maxresults, err := strconv.Atoi(t.Twitter["maxresults"])
	if err != nil {
		fmt.Printf("Reading the twitter max results was a problem\n")
		panic(err)
	}
	config.twitter.maxresults = int64(maxresults)
}

func (config *Config) GetTwitterConsumerKey() string {
	return config.twitter.consumerkey
}

func (config *Config) GetTwitterConsumerSecret() string {
	return config.twitter.consumersecret
}

func (config *Config) GetTwitterAccessToken() string {
	return config.twitter.accesstoken
}

func (config *Config) GetTwitterAccessSecret() string {
	return config.twitter.accesssecret
}

func (config *Config) GetTwitterMaxResultCount() int64 {
	return config.twitter.maxresults
}