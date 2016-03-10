package main

import (
	"fmt"
	"munch/config"
)

func main() {
	cfg := config.New()
	fmt.Printf("Port number: %d\n", cfg.GetWebappPort())
}
