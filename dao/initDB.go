package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := username + ":" + password + "@tcp/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Comment{})
	//db.AutoMigrate(&Follow{})
	db.AutoMigrate(&Video{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}