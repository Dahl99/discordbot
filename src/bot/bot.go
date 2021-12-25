package bot

import (
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
}

func Stop() {
	context.Dg.Close()
}
