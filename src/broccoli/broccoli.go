package main

import "fmt"
import "munch/config"
import "munch/messaging"

func main() {
	cfg := config.New()
	broker := messaging.New(cfg)
	broker.EnqueueCrawl(3)
	fmt.Printf("Use this tool to control munch\n")
}
