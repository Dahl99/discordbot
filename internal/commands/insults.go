package commands

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Dahl99/discordbot/internal/discord"

	"github.com/bwmarrin/discordgo"
)

// insultUrl contains the url for the API generating insults.
const insultUrl string = "https://evilinsult.com/generate_insult.php?lang=en&type=json"

// evilInsultNotAvailable String to be sent if Evil Insult API isn't available.
const evilInsultNotAvailable string = "Evil Insult API not available at the moment. Please try again later."

// Struct to store fetched data from Evil Insult API.
type insult struct {
	Insult string `json:"insult"`
}

func PostInsult(m *discordgo.MessageCreate) {
	discord.SendChannelMessage(m.ChannelID, getInsult())
}

func getInsult() string {
	res, err := http.Get(insultUrl)
	if err != nil {
		slog.Warn("failed to get insult", "error", err)
		return evilInsultNotAvailable
	}

	var insultObj insult

	err = json.NewDecoder(res.Body).Decode(&insultObj) // Decoding data into struct object
	if err != nil {
		slog.Warn("failed to decode response from insult API", "error", err)
		return "ERR: decoding data failed"
	}

	res.Body.Close()

	return insultObj.Insult
}
