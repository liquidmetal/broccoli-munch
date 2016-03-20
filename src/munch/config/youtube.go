package config

import "strconv"
import "fmt"

func (config *Config) parseConfigYoutubeValues(t *ConfigReader) {
	// Mail settings
	config.youtube.clientid = t.Youtube["clientid"]
	config.youtube.clientsecret = t.Youtube["clientsecret"]
	config.youtube.refresh = t.Youtube["refresh"]

	oauthport, err := strconv.Atoi(t.Youtube["oauthport"])
	if err != nil {
		fmt.Printf("Reading the port was a problem\n")
		panic(err)
	}
	config.youtube.oauthport = oauthport

	maxresults, err := strconv.Atoi(t.Youtube["maxresults"])
	if err != nil {
		fmt.Printf("Reading the youtube max results was a problem\n")
		panic(err)
	}
	config.youtube.maxresults = int64(maxresults)
}

func (config *Config) GetYoutubeClientId() string {
	return config.youtube.clientid
}

func (config *Config) GetYoutubeClientSecret() string {
	return config.youtube.clientsecret
}

func (config *Config) GetYoutubeRefreshToken() string {
	return config.youtube.refresh
}

func (config *Config) GetYoutubeOAuthPort() int {
	return config.youtube.oauthport
}

func (config *Config) GetYoutubeMaxResultCount() int64 {
	return config.youtube.maxresults
}
