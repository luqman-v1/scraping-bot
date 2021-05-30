package main

import (
	"context"
	"log"
	"net/http"
	"scraping/cron"
	"scraping/mongodb"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	ctx := context.Background()
	_ = mongodb.Conn(ctx)
	cron.RunJob(ctx)
	_ = http.ListenAndServe("0.0.0.0:8443", nil)
	if err != nil {
		return
	}
}
