package migrations

import (
	"github.com/Dahl99/DiscordBot/internal/database"
	"github.com/Dahl99/DiscordBot/internal/models"
)

func AutoMigrate() {
	database.DB.AutoMigrate(models.ChessGame{})
}
