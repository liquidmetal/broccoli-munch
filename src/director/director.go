package main

import (
	"fmt"
	"munch/config"
	"munch/messaging"
)

func main() {
	//mq := connect()
	//declareQueues(mq)

	cfg := config.New()
	messaging.New(cfg)

	fmt.Printf("This is the munch director. The one who orchestrates everything.\n")
}
