package config

func (config *Config) parseConfigYoutubeValues(t *ConfigReader) {
	// Mail settings
	config.youtube.clientid = t.Youtube["clientid"]
	config.youtube.clientsecret = t.Youtube["clientsecret"]
	config.youtube.refresh = t.Youtube["refresh"]
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
