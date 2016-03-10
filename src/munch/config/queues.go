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
	return nil
}
