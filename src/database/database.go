package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"discordbot/src/models"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/discordbot"), &gorm.Config{})
	if err != nil {
		log.Fatalln("ERR: Could not connect with the database!")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.ChessGame{})
}
