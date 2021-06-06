package action

import (
	"fmt"
	"net/http"
	"scraping/entity"
	mongodb2 "scraping/repo/mongodb"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

func Store(c *gin.Context) {
	var json entity.Link
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.UUID = node.Generate().String()
	_ = mongodb2.Insert(c, json, entity.Links)
	c.JSON(200, gin.H{
		"data": json,
	})
}
