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

	// Register the MessageCreate func as a callback for MessageCreate events.
	dg.AddHandler(discordbot.MessageCreate)

	// Bot needs information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates


	if err := dg.Open(); err != nil { // Creating a connection
		log.Println("Error opening connection,", err)
		return
	}


	// Updating game status of bot
	dg.UpdateGameStatus(1, bot.Status)

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
