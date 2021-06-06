package mongodb

import (
	"context"
	"log"
)

const Database = "scraping_db"

func Insert(ctx context.Context, data interface{}, collectionName string) error {
	collection := ClientMongo.Database(Database).Collection(collectionName)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func InsertMany(ctx context.Context, data []interface{}, collectionName string) error {
	collection := ClientMongo.Database(Database).Collection(collectionName)
	_, err := collection.InsertMany(ctx, data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Find(ctx context.Context, data interface{}, output interface{}, collectionName string) interface{} {
	collection := ClientMongo.Database(Database).Collection(collectionName)
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

func Delete(ctx context.Context, data interface{}, collectionName string) error {
	collection := ClientMongo.Database(Database).Collection(collectionName)
	_, err := collection.DeleteOne(ctx, data)
	if err != nil {
		log.Println("Error at delete mongodb", err)
		return err
	}
	return nil
}
