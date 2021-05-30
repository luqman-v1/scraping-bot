package cron

import (
	"context"
	"log"
	"scraping/tokopedia"
	"time"

	"github.com/robfig/cron/v3"
)

const TimeZone = "Asia/Jakarta"
const CRON = "10 * * * *"

// RunJob process to execute cron job
func RunJob(ctx context.Context) {
	log.Println("Running Job absent ...")

	loc, _ := time.LoadLocation(TimeZone)
	c := cron.New(cron.WithLocation(loc))

	_, err := c.AddFunc(CRON, func() {
		tokopedia.SendNotif(ctx)
	})
	if err != nil {
		log.Println("err", err)
	}
	c.Start()
}
