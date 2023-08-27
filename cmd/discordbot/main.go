package main

import (
	"fmt"
	"github.com/Dahl99/DiscordBot/internal/bot"
	"github.com/Dahl99/DiscordBot/internal/config"
	"github.com/Dahl99/DiscordBot/internal/database"
	"github.com/Dahl99/DiscordBot/internal/database/migrations"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load environment variables
	config.Load()
	if config.IsAppEnvironment(config.APP_ENVIRONMENT_TEST) {
		fmt.Println("App environment is test, aborting startup")
		return
	}

	// Connect to database and run migrations
	database.Connect()
	migrations.AutoMigrate()

	// Start the bot
	bot.Start()

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Stop()
}
