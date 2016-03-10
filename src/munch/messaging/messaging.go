package messaging

import (
	"fmt"
	"github.com/streadway/amqp"
	"munch/config"
)

type Broker struct {
	conn    *amqp.Connection // The connection
	channel *amqp.Channel    // Channel to send commands
	cfg     *config.Config   // Keep a link to the config file (for queue names)
}

func New(config *config.Config) *Broker {
	broker := new(Broker)
	broker.cfg = config

	// Connect to the message queue server
	address := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.GetQueueUsername(),
		config.GetQueuePassword(),
		config.GetQueueHost(),
		config.GetQueuePort())

	conn, err := amqp.Dial(address)
	if err != nil {
		fmt.Printf("There was an error connecting to the broker")
		panic(err)
	}
	broker.conn = conn

	// Create a new channel
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	broker.channel = channel

	broker.initializeQueues()
	return broker
}

func (broker *Broker) initializeQueues() {
	// Setup queues based on the config file
}
