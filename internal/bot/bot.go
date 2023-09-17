package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dahl99/discordbot/internal/commands/chess"
	"github.com/Dahl99/discordbot/internal/commands/music"
	"github.com/Dahl99/discordbot/internal/config"
	"github.com/Dahl99/discordbot/internal/discord"
	"github.com/Dahl99/discordbot/internal/handlers"
)

func Start() {
	config.Load()
	if config.IsAppEnvironment(config.APP_ENVIRONMENT_TEST) {
		fmt.Println("App environment is test, aborting startup")
		return
	}

	discord.InitSession()
	addHandlers()
	discord.InitConnection()

	music.StartRoutine()
	chess.InitAi()

	// Connect to database and run migrations
	//database.Connect()
	//migrations.AutoMigrate()

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	chess.StopChessAi()
	discord.Session.Close()
}

func addHandlers() {
	// Register handlers as callbacks for the events.
	discord.Session.AddHandler(handlers.ReadyHandler)
	// Session.AddHandler(handlers.GuildCreateHandler)
	discord.Session.AddHandler(handlers.MessageCreateHandler)
}
