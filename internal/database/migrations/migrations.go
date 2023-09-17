package migrations

import (
	"github.com/Dahl99/discordbot/internal/database"
	"github.com/Dahl99/discordbot/internal/models"
)

func AutoMigrate() {
	database.DB.AutoMigrate(models.ChessGame{})
}
