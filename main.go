package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"discordbot/src/bot"
)

func main() {

	bot.Start()

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Stop()
}
