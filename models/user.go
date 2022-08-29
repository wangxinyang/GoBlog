package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(50);not null" json:"password"`
	Role     int    `gorm:"type:int;Default:2" json:"role"`
}
