package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"discordbot/src/bot"
	"discordbot/src/config"
	"discordbot/src/database"
)

func main() {

	config.Load()
	if config.IsAppEnvironment(config.APP_ENVIRONMENT_TEST) {
		fmt.Println("App environment is test, aborting startup")
		return
	}

	database.Connect()
	database.AutoMigrate()

	bot.Start()

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Stop()
}
