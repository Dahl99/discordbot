package discordbot

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type kanyeQuote struct {
	Quote string `json:"quote"`
}

func postKanyeQuote(m *discordgo.MessageCreate) {
	dg.ChannelMessageSend(m.ChannelID, getKanyeQuote())
}

func getKanyeQuote() string {
	res, err := http.Get(kanyeRestEndpoint)
	if err != nil {
		log.Println("ERROR: kanye rest API get request failed")
		return kanyeRestUnavailable
	}

	var kanyeQuoteObj kanyeQuote

	err = json.NewDecoder(res.Body).Decode(&kanyeQuoteObj)
	if err != nil {
		log.Println("ERROR: decoding of kanye quote failed")
		return kanyeRestUnavailable
	}

	res.Body.Close()

	return kanyeQuoteObj.Quote
}