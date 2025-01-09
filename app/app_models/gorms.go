package app_models

import (
	"gorm.io/gorm"
)

type LocNames struct {
	gorm.Model
	Id      int    `gorm:"primaryKey, autoIncrement"`
	Locname string `gorm:"size:50" json:"locname"`
}

type ManNames struct {
	gorm.Model
	Id      int    `gorm:"primaryKey, autoIncrement"`
	Manname string `gorm:"size:50" json:"manname"`
}

type TypNames struct {
	gorm.Model
	Id      int    `gorm:"primaryKey, autoIncrement"`
	Typname string `gorm:"size:50" json:"typname"`
}

type StaNames struct {
	gorm.Model
	Id      int    `gorm:"primaryKey, autoIncrement"`
	Staname string `gorm:"size:50" json:"staname"`
}

type Status_History struct {
	gorm.Model
	Id      int    `gorm:"primaryKey autoIncrement"`
	Itmid   int    `gorm:"foreignKey:ItemsItmid" json:"itmid"`
	Staid   int    `gorm:"foreignKey:StaNamesId" json:"staid"`
	UserId  int    `gorm:"foreignKey:UserId" json:"uid"`
	Comment string `gorm:"size:255" json:"comment"`
}

type Items struct {
	gorm.Model
	Id    int `gorm:"primaryKey autoIncrement"`
	Itmid int `gorm:"default:0" json:"itmid"`

	Description string  `gorm:"size:255" json:"description"`
	Serial      string  `gorm:"size:50" json:"serial"`
	Price       float64 `gorm:"size:20" json:"price"`

	Locid int `gorm:"foreignKey:LocNamesId" json:"locid"`
	Typid int `gorm:"foreignKey:TypNamesId" json:"typid"`
	Manid int `gorm:"foreignKey:ManNamesId" json:"manid"`
	Staid int `gorm:"foreignKey:StaNamesId" json:"staid"`

	UserId int `gorm:"foreignKey:UserId" json:"uid"`
}
