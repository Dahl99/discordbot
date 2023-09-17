package discord

import (
	"log/slog"

	"github.com/Dahl99/discordbot/internal/config"

	"github.com/bwmarrin/discordgo"
)

var Session *discordgo.Session

func InitSession() {
	var err error
	Session, err = discordgo.New("Bot " + config.GetDiscordToken()) // Initializing discord session
	if err != nil {
		slog.Error("failed to create discord session", "error", err)
	}

	Session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentGuildMessageTyping | discordgo.IntentGuildVoiceStates | discordgo.IntentGuilds
}

func InitConnection() {
	if err := Session.Open(); err != nil { // Creating a connection
		slog.Error("failed to create websocket connection to discord", "error", err)
		return
	}
}
