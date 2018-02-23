package main

import (
	"github.com/MilesChou/go-post/db"
	"github.com/MilesChou/go-post/routes"
	"github.com/MilesChou/go-post/entity"
)

func main() {
	gorm := db.Connect()
	defer gorm.Close()

	entity.Migrate()

	server := routes.CreateServer()
	server.Run()
}
