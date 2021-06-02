package main

import (
	"discordbot"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	bot := discordbot.ReadJsonBotData()

	dg, err := discordgo.New("Bot " + bot.Token) // Initializing discord session
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(discordbot.MessageCreate)

	if err := dg.Open(); err != nil { // Creating a connection
		log.Println("Error opening connection,", err)
	}

	dg.UpdateGameStatus(1, bot.Status)

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
