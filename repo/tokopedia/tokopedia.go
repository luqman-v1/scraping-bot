package tokopedia

import (
	"context"
	"log"
	"scraping/entity"
	curl2 "scraping/repo/curl"
	mongodb2 "scraping/repo/mongodb"
	telegram2 "scraping/repo/telegram"

	"github.com/PuerkitoBio/goquery"
	"go.mongodb.org/mongo-driver/bson"
)

type Config struct {
	Uri string
}

func GetProduct(ctx context.Context, url string) []entity.Product {
	res := curl2.Get(url)
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]entity.Product, 0)
	doc.Find(".css-jza1fo").Children().Each(func(i int, sel *goquery.Selection) {
		row := new(entity.Product)
		row.Title = sel.Find(".css-18c4yhp").Text()
		row.URL, _ = sel.Find(".css-7fmtuv").Children().Attr("href")
		row.Image, _ = sel.Find(".css-1dpp4z9").Children().Attr("src")
		row.Price = sel.Find(".css-rhd610").Text()
		row.Location = sel.Find(".css-4pwgpi").Text()

		if row.Title != "" && row.URL != "" {
			rows = append(rows, *row)
		}

	})
	return rows
}

func SendNotif(ctx context.Context, url string) {
	log.Println("Send Notif Trigred")
	rows := GetProduct(ctx, url)
	result := make([]entity.Product, 0)
	rest := mongodb2.Find(ctx, bson.M{}, result, entity.Scrapings)
	products := rest.([]entity.Product)
	m := map[string]entity.Product{}
	for _, product := range products {
		m[product.Title+product.Price] = product
	}
	for _, row := range rows {
		_, ok := m[row.Title+row.Price]
		if !ok {
			message :=
				"🔥" + "*" + row.Title + "*" + "\n" +
					"💵" + row.Price + "\n" +
					"🔗" + row.URL + "\n" +
					"📍" + row.Location
			err := telegram2.Send(message)
			if err != nil {
				return
			}
			_ = mongodb2.Insert(ctx, row, entity.Scrapings)
		}
	}
}
