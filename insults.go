package discordbot

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// Struct to store fetched data from Evil Insult API
type insult struct {
	Insult string `json:"insult"`
}

func postInsult(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) == 1 { // Checks if insult command was used properly
		dg.ChannelMessageSend(m.ChannelID, getInsult())
	}
}

func getInsult() string {
	res, err := http.Get(insultURL) // Fetching an insult
	if err != nil {                 // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return evilInsultNotAvailable
	}

	var insultObj insult

	err = json.NewDecoder(res.Body).Decode(&insultObj) // Decoding data into struct object
	if err != nil {
		log.Println(err)
		return decodingFailed
	}

	res.Body.Close() // Closing body to prevent resource leak

	return insultObj.Insult
}
