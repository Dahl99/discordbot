package database

import (
	"github.com/Dahl99/DiscordBot/internal/config"
	"log/slog"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	for i := 0; i < 4; i++ {
		dsn := config.GetDatabaseUsername() + ":" +
			config.GetDatabasePassword() + "@tcp(mysql:" +
			config.GetDatabasePort() + ")/" +
			config.GetDatabaseDatabase() + "?parseTime=true"

		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			slog.Warn("failed to connect to database, retrying in 5 seconds", "error", err)
			time.Sleep(time.Second * 5)
		}
	}

	if err != nil {
		slog.Error("failed to establish connection to database", "error", err)
	}
}
