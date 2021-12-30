package bot

import (
	"discordbot/src/commands/chess"
	"discordbot/src/config"
	"discordbot/src/context"
	"discordbot/src/handlers"
	"discordbot/src/music"
	"math/rand"
	"time"
)

func Start() {
	rand.Seed(time.Now().UnixNano())
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
