package main

import (
	"discordbot"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	discordbot.Initialize()

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discordbot.SafeDestroy()
}
