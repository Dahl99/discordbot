package bot

import (
	"discordbot/src/handlers"
	"discordbot/src/music"
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

// Global struct object containing bot config
var Config *configuration
var Dg *discordgo.Session

// readconfigig reads the data the bot needs from the provided JSON file
func readConfiguration() {
	res, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(res, &Config)
	if err != nil {
		log.Println(err)
	}
}

func Initialize() {
	readConfiguration()
	
	var err error
	Dg, err = discordgo.New("Bot " + Config.Token) // Initializing discord session
	if err != nil {
		log.Println("ERROR: error creating Discord session,", err)
		return
	}

	// Register handlers as callbacks for the events.
	Dg.AddHandler(handlers.ReadyHandler)
	Dg.AddHandler(handlers.GuildCreateHandler)
	Dg.AddHandler(handlers.MessageCreateHandler)

	if err := Dg.Open(); err != nil { // Creating a connection
		log.Println("ERROR: unable to open connection,", err)
		return
	}

	initializeRoutine()
}

func initializeRoutine() {
	music.SongSignal = make(chan music.PkgSong)
	go music.GlobalPlay(music.SongSignal)
}

func SafeDestroy() {
	Dg.Close()
}
