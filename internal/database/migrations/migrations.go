package migrations

import (
	"github.com/Dahl99/discord-bot/internal/database"
	"github.com/Dahl99/discord-bot/internal/models"
)

func AutoMigrate() {
	database.DB.AutoMigrate(models.ChessGame{})
}
