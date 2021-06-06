package main

import (
	"context"
	"log"
	cron2 "scraping/repo/cron"
	mongodb2 "scraping/repo/mongodb"
	"scraping/route"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	ctx := context.Background()
	_ = mongodb2.Conn(ctx)
	cron2.RunJob(ctx)
	route.Run()
}
