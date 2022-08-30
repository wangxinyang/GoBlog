package routers

import (
	"github.com/gin-gonic/gin"
	v1 "goblog/api/v1"
	"goblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("/api/v1")

	// router group
	{
		// ユーザーモジュールのルーター
		router.POST("/user/add", v1.AddUser)
		router.GET("/users", v1.GetUsers)
		router.PUT("/user/:id", v1.EditUser)
		router.DELETE("/user/:id", v1.DeleteUser)

		// カテゴリモジュールのルーター

		// アーティクルモジュールのルーター

	}
	r.Run(utils.HttpPort)
}
