package entity

type Link struct {
	UUID        string `bson:"uuid" json:"uuid"`
	Url         string `bson:"url" json:"url"`
	MarketPlace string `bson:"market_place" json:"market_place"`
}
