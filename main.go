package main

import (
	"goblog/models"
	"goblog/routers"
)

func main() {
	// データベース初期化
	models.InitDb()

	// ルーター初期化
	routers.InitRouter()
}
