package entity

import "github.com/MilesChou/go-post/db"

func Migrate() {
	gorm := db.Connect()
	gorm.AutoMigrate(&Post{})
}
