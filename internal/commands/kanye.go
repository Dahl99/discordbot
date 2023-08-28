package commands

import (
	"encoding/json"
	"github.com/Dahl99/discord-bot/internal/discord"
	"log/slog"
	"net/http"

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
	discord.SendChannelMessage(m.ChannelID, getKanyeQuote())
}

func getKanyeQuote() string {
	res, err := http.Get(kanyeRestEndpoint)
	if err != nil {
		slog.Warn("failed to get Kanye quote", "error", err)
		return kanyeRestUnavailable
	}

	var kanyeQuoteObj kanyeQuote

	err = json.NewDecoder(res.Body).Decode(&kanyeQuoteObj)
	if err != nil {
		slog.Warn("failed to decode response from Kanye API", "error", err)
		return "ERR: decoding data failed"
	}

	res.Body.Close()

	return kanyeQuoteObj.Quote
}
