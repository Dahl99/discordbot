package discordbot

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
)

type config struct {
	Token string
	Prefix string
	Status string
	Online string
	Ytkey string
}

// Global struct object containing bot config
var conf config
var dg *discordgo.Session

// readConfig reads the data the bot needs from the provided JSON file
func readConfig() {
	res, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(res, &conf)
	if err != nil {
		log.Println(err)
	}
}

func Initialize() {
	readConfig()
	
	var err error
	dg, err = discordgo.New("Bot " + conf.Token) // Initializing discord session
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(ReadyHandler)

	// Register the MessageCreate func as a callback for MessageCreate events.
	dg.AddHandler(MessageCreateHandler)

	// Register guildCreate as a callback for the guildCreate events.
	// dg.AddHandler(GuildCreateHandler)

	// Bot needs information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates


	if err := dg.Open(); err != nil { // Creating a connection
		log.Println("Error opening connection,", err)
		return
	}
}

func SafeDestroy() {
	dg.Close()
}
