package commands

import (
	"encoding/json"
	"log"
	"net/http"

	"discordbot/src/utils"

	"github.com/bwmarrin/discordgo"
)

// kanyeRestEndpoint contains endpoint for kanye rest API
const kanyeRestEndpoint string = "https://api.kanye.rest/"

// kanyeRestUnavailable is to be sent if kanye rest api is unavailable
const kanyeRestUnavailable string = "Oops, something went wrong when getting Kanye quote"

type kanyeQuote struct {
	Quote string `json:"quote"`
}

func PostKanyeQuote(m *discordgo.MessageCreate) {
	utils.SendChannelMessage(m.ChannelID, getKanyeQuote())
}

func getKanyeQuote() string {
	res, err := http.Get(kanyeRestEndpoint)
	if err != nil {
		log.Println(err)
		return kanyeRestUnavailable
	}

	var kanyeQuoteObj kanyeQuote

	err = json.NewDecoder(res.Body).Decode(&kanyeQuoteObj)
	if err != nil {
		log.Println(err)
		return "ERR: decoding data failed"
	}

	res.Body.Close()

	return kanyeQuoteObj.Quote
}
