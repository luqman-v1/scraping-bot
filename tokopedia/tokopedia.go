package tokopedia

import (
	"context"
	"log"
	"os"
	"scraping/curl"
	"scraping/entity"
	"scraping/mongodb"
	"scraping/telegram"

	"github.com/PuerkitoBio/goquery"
	"go.mongodb.org/mongo-driver/bson"
)

type Config struct {
	Uri string
}

func GetProduct(ctx context.Context) []entity.Product {
	url := os.Getenv("PRODUCT_URL")
	res := curl.Get(url)
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

func SendNotif(ctx context.Context) {
	log.Println("Send Notif Trigred")
	rows := GetProduct(ctx)
	result := make([]entity.Product, 0)
	rest := mongodb.Find(ctx, bson.M{}, result)
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
			err := telegram.Send(message)
			if err != nil {
				return
			}
			_ = mongodb.Insert(ctx, row)
		}
	}
}
