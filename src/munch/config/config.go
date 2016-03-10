package config

import (
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"strconv"
)

type ConfigReader struct {
	Mail   map[string]string
	Db     map[string]string
	Webapp map[string]string
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
	// Mail settings
	config.mail.publickey = t.Mail["publickey"]
	config.mail.privatekey = t.Mail["privatekey"]

	// Database settings
	config.db.filename = t.Db["filename"]

	// Webapp settings
	pnum, err := strconv.Atoi(t.Webapp["port"])
	if err != nil {
		fmt.Errorf("There was an error parsing the webapp port number\n")
		return err
	}
	config.webapp.port = pnum

	return nil
}

func (config *Config) GetMailPublicKey() string {
	return config.mail.publickey
}

func (config *Config) GetMailPrivateKey() string {
	return config.mail.privatekey
}

func (config *Config) GetDbFilename() string {
	return config.db.filename
}

func (config *Config) GetWebappPort() int {
	return config.webapp.port
}
