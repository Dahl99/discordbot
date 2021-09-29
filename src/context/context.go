package context

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var Dg *discordgo.Session

func Initialize(discordToken string) {
	var err error
	Dg, err = discordgo.New("Bot " + discordToken) // Initializing discord session
	if err != nil {
		log.Fatalln("ERROR: error creating Discord session,", err)
		return
	}
}

func OpenConnection() {
	if err := Dg.Open(); err != nil { // Creating a connection
		log.Fatalln("ERROR: unable to open connection,", err)
		return
	}
}
