package bot

import (
	"discordbot/src/config"
	"discordbot/src/handlers"
	"discordbot/src/music"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Initialize() {
	config.ReadConfiguration()
	
	var err error
	config.Dg, err = discordgo.New("Bot " + config.Config.Token) // Initializing discord session
	if err != nil {
		log.Println("ERROR: error creating Discord session,", err)
		return
	}

	// Register handlers as callbacks for the events.
	config.Dg.AddHandler(handlers.ReadyHandler)
	config.Dg.AddHandler(handlers.GuildCreateHandler)
	config.Dg.AddHandler(handlers.MessageCreateHandler)

	if err := config.Dg.Open(); err != nil { // Creating a connection
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
	config.Dg.Close()
}
