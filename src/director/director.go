package main

import (
	"fmt"
	"munch/config"
	"munch/messaging"
)

func main() {

	cfg := config.New()
	messaging.New(cfg)

	fmt.Printf("This is the munch director. The one who orchestrates everything.\n")
}
