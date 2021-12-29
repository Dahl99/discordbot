package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/discordbot?parseTime=true"), &gorm.Config{})
		if err != nil {
			log.Println("INFO: Could not connect with the database, retrying in 2 seconds")
			time.Sleep(time.Second * 2)
		}
	}

	if err != nil {
		log.Fatalln("ERR: Could not connect with the database!")
	}
}
