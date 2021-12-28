package migrations

import (
	"discordbot/src/database"
	"discordbot/src/models"
)

func AutoMigrate() {
	database.DB.AutoMigrate(models.ChessGame{})
}
