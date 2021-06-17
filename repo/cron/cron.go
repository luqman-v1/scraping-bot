package cron

import (
	"context"
	"log"
	"scraping/entity"
	mongodb2 "scraping/repo/mongodb"
	tokopedia2 "scraping/repo/tokopedia"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/robfig/cron/v3"
)

const TimeZone = "Asia/Jakarta"
const CRON = "*/10 * * * *"

// RunJob process to execute cron job
func RunJob(ctx context.Context) {
	log.Println("Running Job absent ...")

	loc, _ := time.LoadLocation(TimeZone)
	c := cron.New(cron.WithLocation(loc))

	_, err := c.AddFunc(CRON, func() {
		result := make([]entity.Link, 0)
		rest := mongodb2.Find(ctx, bson.M{}, result, entity.Links)
		links := rest.([]entity.Link)
		for _, link := range links {
			tokopedia2.SendNotif(ctx, link.Url)
		}
	})
	if err != nil {
		log.Println("err", err)
	}
	c.Start()
}
