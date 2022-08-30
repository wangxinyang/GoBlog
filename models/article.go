package models

import (
	"goblog/code"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	gorm.Model
	Title        string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int      `gorm:"type:int;not null" json:"cid"`
	Desc         string   `gorm:"type:varchar(200);" json:"desc"`
	Content      string   `gorm:"type:longtext;" json:"content"`
	Img          string   `gorm:"type:varchar(100);" json:"img"`
	CommentCount int      `gorm:"type:int;not null;Default:0" json:"comment_count"`
	ReadCount    int      `gorm:"type:int;not null;Default:0" json:"read_count"`
	Category     Category `gorm:"foreignkey:Cid"`
}

// CreateArticle create article
func CreateArticle(data Article) int {
	err := Db.Create(&data).Error
	if err != nil {
		log.Printf("create article was error happened: %s", err)
		return code.ERROR
	}
	return code.SUCCESS
}

// GetArticleList アーティクルリストを取得する
func GetArticleList(page int, size int) (int, []Article) {
	var article []Article
	err := Db.Preload("Category").Offset(page).Limit(size).Find(&article).Error
	if err != nil {
		log.Printf("query category list was error happened: %s", err)
		return code.ERROR, article
	}
	return code.SUCCESS, article
}

// EditArticle edit article by id
func EditArticle(id int, article Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["desc"] = article.Desc
	maps["content"] = article.Content
	maps["img"] = article.Img

	err := Db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

// DeleteArticleById delete article by id
func DeleteArticleById(id int) int {
	var article Article
	err := Db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}
