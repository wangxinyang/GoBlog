package models

import (
	"goblog/code"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(50);not null" json:"password"`
	Role     int    `gorm:"type:int;Default:2" json:"role"`
}

// CheckUserExists check user exists
func CheckUserExists(username string) int {
	var user User
	var count int64
	err := Db.Model(&user).Where("username = ?", username).Count(&count).Error
	if err != nil {
		log.Printf("query user table was error happened: %s", err)
		return code.ERROR
	}
	if count > 0 {
		log.Printf("user was existed")
		return code.ERROR_USERNAME_USED
	}
	return code.SUCCESS
}

// CreateUser create user
func CreateUser(data User) int {
	err := Db.Create(&data).Error
	if err != nil {
		log.Printf("create user was error happened: %s", err)
		return code.ERROR
	}
	return code.SUCCESS
}

// GetUserList ユーザーリストを取得する
func GetUserList(page int, size int) (int, []User) {
	var user []User
	err := Db.Offset(page).Limit(size).Find(&user).Error
	if err != nil {
		log.Printf("query user list was error happened: %s", err)
		return code.ERROR, user
	}
	return code.SUCCESS, user
}

func EditUser(id int, user User) int {
	var maps = make(map[string]interface{})
	maps["username"] = user.Username
	maps["role"] = user.Role

	err := Db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

// DeleteUserById delete user by id
func DeleteUserById(id int) int {
	var user User
	err := Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}
