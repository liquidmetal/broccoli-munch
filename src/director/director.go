package main

import (
	"fmt"
	"munch/config"
	"munch/database"
	"munch/messaging"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	cfg := config.New()
	db := database.New(cfg)
	broker := messaging.New(cfg)

	quit := make(chan os.Signal, 1)
	quit_crawl := make(chan os.Signal, 1)
	quit_email := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	go func() {
		for {
			select {
			case c := <-quit:
				quit_crawl <- c
				quit_email <- c
			}
		}
	}()

	wg := new(sync.WaitGroup)

	// Every n minutes
	num_minutes_crawl := 1 * time.Minute
	ticker_crawl := time.NewTicker(num_minutes_crawl)
	go handleCrawls(cfg, db, broker, ticker_crawl, quit_crawl, wg)

	// Every n minutes
	num_minutes_email := 24 * 60 * time.Minute
	ticker_email := time.NewTicker(num_minutes_email)
	go handleEmails(cfg, db, broker, ticker_email, quit_email, wg)
	wg.Add(2)

	wg.Wait()

	// Cleanup
	db.Close()
}

func handleCrawls(cfg *config.Config, db *database.Db, broker *messaging.Broker, ticker *time.Ticker, quit chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ticker.C:
			fmt.Printf("Enqueueing crawl requests ..\n")
			ids, err := db.FetchAllSourceIds()
			if err != nil {
				continue
			}
			for _, id := range ids {
				broker.EnqueueCrawl(int(id))
			}
			fmt.Printf("Enqueueing crawl requests [done]\n")

		case <-quit:
			fmt.Printf("Shutting down crawling loop\n")
			ticker.Stop()
			return
		}
	}
}

func handleEmails(cfg *config.Config, db *database.Db, broker *messaging.Broker, ticker *time.Ticker, quit chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ticker.C:
			fmt.Printf("Email stuff\n")

		case <-quit:
			fmt.Printf("Shutting down email loop\n")
			ticker.Stop()
			return
		}
	}
}
