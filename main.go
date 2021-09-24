package main

import (
	"discordbot/src/bot"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	bot.Initialize()

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.SafeDestroy()
}
