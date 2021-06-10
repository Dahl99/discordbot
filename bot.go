package discordbot

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
)

type configuration struct {
	Token string
	Prefix string
	Status string
	Online string
	Ytkey string
}

// Global struct object containing bot configig
var config *configuration
var dg *discordgo.Session

// readconfigig reads the data the bot needs from the provided JSON file
func readConfiguration() {
	res, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(res, &config)
	if err != nil {
		log.Println(err)
	}
}

func Initialize() {
	readConfiguration()
	var err error
	dg, err = discordgo.New("Bot " + config.Token) // Initializing discord session
	if err != nil {
		log.Println("ERROR: error creating Discord session,", err)
		return
	}

	// Register handlers as callbacks for the events.
	dg.AddHandler(ReadyHandler)
	dg.AddHandler(GuildCreateHandler)
	dg.AddHandler(MessageCreateHandler)

	if err := dg.Open(); err != nil { // Creating a connection
		log.Println("Error opening connection,", err)
		return
	}

	initializeRoutine()
}

func initializeRoutine() {
	songSignal = make(chan PkgSong)
	go GlobalPlay(songSignal)
}

func SafeDestroy() {
	dg.Close()
}
