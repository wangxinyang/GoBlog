package models

import (
	"fmt"
	"goblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var Db = InitDb()

// InitDb データベースに繋がる
func InitDb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println("データベースに繋がった時エラーになります。", err)
	}

	// 移行
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("ジェネレーターDBオブジェクトを取得した時エラーになりました。", err)
	}

	// SetMaxIdleConns コンネクションプールに最大のコンネクション数
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 最大DBコンネクション数
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 最大lifetime。
	sqlDB.SetConnMaxLifetime(5 * time.Second)

	return db
}
