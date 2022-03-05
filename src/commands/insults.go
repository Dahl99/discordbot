package commands

import (
	"encoding/json"
	"log"
	"net/http"

	"discordbot/src/utils"

	"github.com/getsentry/sentry-go"

	"github.com/bwmarrin/discordgo"
)

// insultUrl contains the url for the API generating insults
const insultUrl string = "https://evilinsult.com/generate_insult.php?lang=en&type=json"

// evilInsultNotAvailable String to be sent if Evil Insult API isn't available
const evilInsultNotAvailable string = "Evil Insult API not available at the moment. Please try again later."

// Struct to store fetched data from Evil Insult API
type insult struct {
	Insult string `json:"insult"`
}

func PostInsult(m *discordgo.MessageCreate) {
	utils.SendChannelMessage(m.ChannelID, getInsult())
}

func getInsult() string {
	res, err := http.Get(insultUrl) // Fetching an insult
	if err != nil {
		sentry.CaptureException(err) // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return evilInsultNotAvailable
	}

	var insultObj insult

	err = json.NewDecoder(res.Body).Decode(&insultObj) // Decoding data into struct object
	if err != nil {
		sentry.CaptureException(err)
		log.Println(err)
		return "ERR: decoding data failed"
	}

	res.Body.Close() // Closing body to prevent resource leak

	return insultObj.Insult
}
