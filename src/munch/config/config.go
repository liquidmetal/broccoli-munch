package config

import (
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type ConfigReader struct {
	Mail map[string]string
}

type Config struct {
	mail struct {
		publickey  string `yaml:"publickey"`
		privatekey string `yaml:"privatekey"`
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

	config.parseConfigValues(t)

	return nil
}

func (config *Config) parseConfigValues(t *ConfigReader) error {
	config.mail.publickey = t.Mail["publickey"]
	config.mail.privatekey = t.Mail["privatekey"]

	return nil
}

func (config *Config) GetMailPublicKey() string {
	return config.mail.publickey
}

func (config *Config) GetMailPrivateKey() string {
	return config.mail.privatekey
}
