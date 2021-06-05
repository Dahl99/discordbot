package discordbot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

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
		log.Println("ERROR: error creating Discord session,", err)
		return
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(ReadyHandler)

	// Register the MessageCreate func as a callback for MessageCreate events.
	dg.AddHandler(MessageCreateHandler)

	// Register guildCreate as a callback for the guildCreate events.
	// dg.AddHandler(GuildJoinHandler)

	// Bot needs information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates


	if err := dg.Open(); err != nil { // Creating a connection
		log.Println("Error opening connection,", err)
		return
	}

	// purgeRoutine()
	initializeRoutine()
}

// func purgeRoutine() {
// 	go func() {
// 		for {
// 			for k, v := range purgeQueue {
// 				if time.Now().Unix()-o.DiscordPurgeTime > v.TimeSent {
// 					purgeQueue = append(purgeQueue[:k], purgeQueue[k+1:]...)
// 					dg.ChannelMessageDelete(v.ChannelID, v.ID)
// 					break
// 				}
// 			}
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()
// }

func initializeRoutine() {
	songSignal = make(chan PkgSong)
	go GlobalPlay(songSignal)
}

func GlobalPlay(songSig chan PkgSong) {
	for {
		select {
		case song := <-songSig:
			if song.v.radioFlag {
				song.v.Stop()
				time.Sleep(200 * time.Millisecond)
			}
			go song.v.PlayQueue(song.data)
		}
	}
}

func SafeDestroy() {
	dg.Close()
}
