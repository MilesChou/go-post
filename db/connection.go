package db

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const FILE_PATH = "./posts.db"

var instance *gorm.DB

func Connect() *gorm.DB {
	os.Chdir(".")
	if instance == nil {
		db, err := gorm.Open("sqlite3", FILE_PATH)

		if err != nil {
			panic(err.Error())
		}

		instance = db
	}

	return instance
}

func Close() {
	if instance != nil {
		instance.Close()
	}
}
