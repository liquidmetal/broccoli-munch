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
