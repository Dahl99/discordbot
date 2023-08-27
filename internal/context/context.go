package context

import (
	"github.com/Dahl99/DiscordBot/internal/config"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var Dg *discordgo.Session

func Initialize() {
	var err error
	Dg, err = discordgo.New("Bot " + config.GetDiscordToken()) // Initializing discord session
	if err != nil {
		slog.Error("failed to create discord session", "error", err)
		return
	}
}

func OpenConnection() {
	if err := Dg.Open(); err != nil { // Creating a connection
		slog.Error("failed to create websocket connection to discord", "error", err)
		return
	}
}
