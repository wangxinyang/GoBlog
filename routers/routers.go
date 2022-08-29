package routers

import (
	"github.com/gin-gonic/gin"
	"goblog/utils"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	v1 := r.Group("/api/v1")

	// router group
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello, world!")
		})
	}
	r.Run(utils.HttpPort)
}
