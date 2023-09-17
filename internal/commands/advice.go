package commands

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Dahl99/discordbot/internal/discord"

	"github.com/bwmarrin/discordgo"
)

// adviceSlipURL contains url to adviceslip API.
const adviceSlipURL string = "https://api.adviceslip.com/advice"

// adviceSlipNotAvailable contains string to be sent if adviceslip API is unavailable.
const adviceSlipNotAvailable string = "Adviceslip API not available at the moment."

// Struct used to store advice in json.
type slip struct {
	Advice string `json:"advice"`
}

// Struct containing the advice slip.
type allSlips struct {
	Slips slip `json:"slip"`
}

func PostAdvice(m *discordgo.MessageCreate) {
	discord.SendChannelMessage(m.ChannelID, getAdvice())
}

func getAdvice() string {
	res, err := http.Get(adviceSlipURL)
	if err != nil {
		slog.Warn("failed to get advice", "error", err)
		return adviceSlipNotAvailable
	}

	var slips allSlips
	err = json.NewDecoder(res.Body).Decode(&slips)
	if err != nil {
		slog.Warn("failed to decode response from advice API", "error", err)
		return "ERR: decoding data failed"
	}

	res.Body.Close()

	return slips.Slips.Advice
}
