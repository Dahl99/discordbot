package bot

import (
	"discordbot/src/commands/chess"
	"discordbot/src/config"
	"discordbot/src/context"
	"discordbot/src/handlers"
	"discordbot/src/music"
)

func Start() {
	context.Initialize(config.GetDiscordToken())
	handlers.AddHandlers()
	context.OpenConnection()
	music.InitializeRoutine()
	chess.InitChessAi()
}

func Stop() {
	chess.StopChessAi()
	context.Dg.Close()
}
