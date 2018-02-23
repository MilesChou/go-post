package entity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/MilesChou/go-post/db"
)

type Post struct {
	gorm.Model
	Name    string
	Content string
}

func CreatePost(name string, content string) Post {
	post := Post{
		Name:    name,
		Content: content,
	}

	gorm := db.Connect()
	gorm.Create(&post)

	return post
}

func ReadPost(id uint) Post {
	var post Post

	gorm := db.Connect()
	gorm.First(&post, id)

	return post
}
