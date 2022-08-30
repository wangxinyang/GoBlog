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
		router.POST("/category/add", v1.AddCategory)
		router.GET("/categories", v1.GetCategoryList)
		router.PUT("/category/:id", v1.EditCategory)
		router.DELETE("/category/:id", v1.DeleteCategory)

		// アーティクルモジュールのルーター
		router.POST("/article/add", v1.AddArticle)
		router.GET("/articles", v1.GetArticleList)
		router.PUT("/article/:id", v1.EditArticle)
		router.DELETE("/article/:id", v1.DeleteArticle)
	}
	r.Run(utils.HttpPort)
}
