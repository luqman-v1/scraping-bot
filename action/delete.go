package action

import (
	"net/http"
	"scraping/entity"
	mongodb2 "scraping/repo/mongodb"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	var selector = bson.M{"uuid": uuid}
	if err := mongodb2.Delete(c, selector, entity.Links); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": "success",
	})
}
