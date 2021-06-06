package action

import (
	"scraping/entity"
	mongodb2 "scraping/repo/mongodb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func List(c *gin.Context) {
	result := make([]entity.Link, 0)
	rest := mongodb2.Find(c, bson.M{}, result, entity.Links)
	c.JSON(200, gin.H{
		"data": rest,
	})
}
