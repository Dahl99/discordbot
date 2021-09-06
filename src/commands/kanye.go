package commands

import (
	"discordbot/src/bot"
	"discordbot/src/consts"
	"encoding/json"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type kanyeQuote struct {
	Quote string `json:"quote"`
}

func PostKanyeQuote(m *discordgo.MessageCreate) {
	bot.Dg.ChannelMessageSend(m.ChannelID, getKanyeQuote())
}

func getKanyeQuote() string {
	res, err := http.Get(consts.KanyeRestEndpoint)
	if err != nil {
		log.Println("ERROR: kanye rest API get request failed")
		return consts.KanyeRestUnavailable
	}

	var kanyeQuoteObj kanyeQuote

	err = json.NewDecoder(res.Body).Decode(&kanyeQuoteObj)
	if err != nil {
		log.Println("ERROR: decoding of kanye quote failed")
		return consts.KanyeRestUnavailable
	}

	res.Body.Close()

	return kanyeQuoteObj.Quote
}