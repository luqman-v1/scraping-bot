package entity

type Product struct {
	Title    string `bson:"title"`
	URL      string `bson:"url"`
	Image    string `bson:"image"`
	Price    string `bson:"price"`
	Location string `bson:"location"`
}
