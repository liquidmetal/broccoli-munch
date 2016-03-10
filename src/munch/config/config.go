package config

import (
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type ConfigReader struct {
	Mail   map[string]string
	Db     map[string]string
	Webapp map[string]string
	Queues map[string]string
}

type Config struct {
	mail struct {
		publickey  string
		privatekey string
	}

	db struct {
		filename string
	}

	webapp struct {
		port int
	}

	queues struct {
		host     string
		port     int
		username string
		password string
	}
}

func New() *Config {
	ret := new(Config)

	ret.readConfig()
	return ret
}

func (config *Config) readConfig() error {
	filename := "config.yaml"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("There was an error reading the config file:\n%s\n", err)
		return errors.New("Unable to read config")
	}

	t := new(ConfigReader)
	err = yaml.Unmarshal(bytes, &t)
	if err != nil {
		fmt.Printf("There was an error reading the config file:\n%s\n", err)
		return errors.New("Unable to read config")
	}

	err = config.parseConfigValues(t)
	if err != nil {
		fmt.Printf("There was a problem parsing the config file\n")
		panic(err)
	}

	return nil
}

func (config *Config) parseConfigValues(t *ConfigReader) error {
	config.parseConfigMailValues(t)

	// Database settings
	config.db.filename = t.Db["filename"]

	err := config.parseConfigWebappValues(t)
	if err != nil {
		panic(err)
	}

	err = config.parseConfigQueueValues(t)
	if err != nil {
		panic(err)
	}

	return nil
}

func (config *Config) GetDbFilename() string {
	return config.db.filename
}
