package bot

import (
	"github.com/Dahl99/DiscordBot/internal/commands/chess"
	"github.com/Dahl99/DiscordBot/internal/context"
	"github.com/Dahl99/DiscordBot/internal/handlers"
	"github.com/Dahl99/DiscordBot/internal/music"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Start() {
	rand.Seed(time.Now().UnixNano())
	context.Initialize()
	handlers.AddHandlers()
	context.Dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentGuildMessageTyping
	context.OpenConnection()
	music.InitializeRoutine()
	chess.InitChessAi()
}

func Stop() {
	chess.StopChessAi()
	context.Dg.Close()
}
