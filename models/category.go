package models

import (
	"goblog/code"
	"log"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategoryExists check category exists
func CheckCategoryExists(name string) int {
	var category Category
	var count int64
	err := Db.Model(&category).Where("name = ?", name).Count(&count).Error
	if err != nil {
		log.Printf("query category table was error happened: %s", err)
		return code.ERROR
	}
	if count > 0 {
		log.Printf("category was existed")
		return code.ERROR_CATEGORYNAME_USED
	}
	return code.SUCCESS
}

// CreateCategory create category
func CreateCategory(data Category) int {
	err := Db.Create(&data).Error
	if err != nil {
		log.Printf("create category was error happened: %s", err)
		return code.ERROR
	}
	return code.SUCCESS
}

// GetCategoryList カテゴリリストを取得する
func GetCategoryList(page int, size int) (int, []Category) {
	var category []Category
	err := Db.Offset(page).Limit(size).Find(&category).Error
	if err != nil {
		log.Printf("query category list was error happened: %s", err)
		return code.ERROR, category
	}
	return code.SUCCESS, category
}

// EditCategory edit category by id
func EditCategory(id int, category Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = category.Name

	err := Db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

// DeleteCategoryById delete category by id
func DeleteCategoryById(id int) int {
	var user User
	err := Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}
