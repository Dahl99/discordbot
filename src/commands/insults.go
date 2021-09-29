package commands

import (
	"encoding/json"
	"log"
	"net/http"

	"discordbot/src/consts"
	"discordbot/src/utils"

	"github.com/bwmarrin/discordgo"
)

// Struct to store fetched data from Evil Insult API
type insult struct {
	Insult string `json:"insult"`
}

func PostInsult(m *discordgo.MessageCreate) {
	utils.SendChannelMessage(m.ChannelID, getInsult())
}

func getInsult() string {
	res, err := http.Get(consts.InsultURL) // Fetching an insult
	if err != nil {                        // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return consts.EvilInsultNotAvailable
	}

	var insultObj insult

	err = json.NewDecoder(res.Body).Decode(&insultObj) // Decoding data into struct object
	if err != nil {
		log.Println(err)
		return consts.DecodingFailed
	}

	res.Body.Close() // Closing body to prevent resource leak

	return insultObj.Insult
}
