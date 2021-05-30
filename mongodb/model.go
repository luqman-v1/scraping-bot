package mongodb

import (
	"context"
	"log"
)

func Insert(ctx context.Context, data interface{}) error {
	collection := ClientMongo.Database("scraping_db").Collection("scrapings")
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func InsertMany(ctx context.Context, data []interface{}) error {
	collection := ClientMongo.Database("scraping_db").Collection("scrapings")
	_, err := collection.InsertMany(ctx, data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Find(ctx context.Context, data interface{}, output interface{}) interface{} {
	collection := ClientMongo.Database("scraping_db").Collection("scrapings")
	find, err := collection.Find(ctx, data)
	if err != nil {
		return nil
	}
	err = find.All(ctx, &output)
	if err != nil {
		return nil
	}
	return output
}
