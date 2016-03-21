// Functions related to message queue settings
package config

import (
	"fmt"
	"strconv"
)

func (config *Config) parseConfigQueueValues(t *ConfigReader) error {
	// Message queue settings
	config.queues.host = t.Queues["host"]
	pnum, err := strconv.Atoi(t.Queues["port"])
	if err != nil {
		fmt.Printf("There was an error parsing the message queue port number\n")
		return err
	}
	config.queues.port = pnum
	config.queues.username = t.Queues["username"]
	config.queues.password = t.Queues["password"]
	config.queues.queue_crawl = t.Queues["queue_crawl"]
	config.queues.exchange = t.Queues["exchange"]
	return nil
}

func (config *Config) GetQueueHost() string {
	return config.queues.host
}

func (config *Config) GetQueuePort() int {
	return config.queues.port
}

func (config *Config) GetQueueUsername() string {
	return config.queues.username
}

func (config *Config) GetQueuePassword() string {
	return config.queues.password
}

func (config *Config) GetQueueCrawl() string {
	return config.queues.queue_crawl
}

func (config *Config) GetQueueExchange() string {
	return config.queues.exchange
}
