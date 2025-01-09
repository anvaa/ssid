package users

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id         int    `gorm:"primaryKey"`
	Email      string `gorm:"unique, size:255" json:"email"`
	Password   string `gorm:"not null, size:255" json:"password"`
	Role       string `gorm:"default:user, size:20" json:"role"`
	IsAuth     bool   `gorm:"default:false" json:"is_auth"`
	AccessTime int    `gorm:"default:43200" json:"accesstime"` // 12 hour
}

type Links struct {
	gorm.Model
	Id     int    `gorm:"primaryKey, autoIncrement"`
	Url    string `gorm:"default:/" json:"url"`
	UserId int    `gorm:"foreignKey:UserId" json:"uid"`
}
