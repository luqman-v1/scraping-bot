package route

import (
	"scraping/action"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/api/list", action.List)
	r.POST("/api/link/store", action.Store)
	r.DELETE("/api/link/delete/:uuid", action.Delete)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
