package config

import (
	"fmt"
	"strconv"
)

func (config *Config) parseConfigWebappValues(t *ConfigReader) error {
	// Webapp settings
	pnum, err := strconv.Atoi(t.Webapp["port"])
	if err != nil {
		fmt.Errorf("There was an error parsing the webapp port number\n")
		return err
	}
	config.webapp.port = pnum
	return nil
}

func (config *Config) GetWebappPort() int {
	return config.webapp.port
}
