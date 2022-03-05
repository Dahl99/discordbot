package commands

import (
	"encoding/json"
	"log"
	"net/http"

	"discordbot/src/utils"

	"github.com/getsentry/sentry-go"

	"github.com/bwmarrin/discordgo"
)

// adviceSlipURL contains url to adviceslip API
const adviceSlipURL string = "https://api.adviceslip.com/advice"

// adviceSlipNotAvailable contains string to be sent if adviceslip API is unavailable
const adviceSlipNotAvailable string = "Adviceslip API not available at the moment."

//Struct used to store advice in json
type slip struct {
	Advice string `json:"advice"`
}

//Struct containing the advice slip
type allSlips struct {
	Slips slip `json:"slip"`
}

func PostAdvice(m *discordgo.MessageCreate) {
	utils.SendChannelMessage(m.ChannelID, getAdvice())
}

func getAdvice() string {
	res, err := http.Get(adviceSlipURL) // Fetching an advice
	if err != nil {
		sentry.CaptureException(err)
		log.Println(http.StatusServiceUnavailable)
		return adviceSlipNotAvailable
	}

	//	Decoding results into autoresult struct object
	var slips allSlips
	err = json.NewDecoder(res.Body).Decode(&slips)
	if err != nil {
		sentry.CaptureException(err)
		log.Println(err)
		return "ERR: decoding data failed"
	}
	res.Body.Close() // Closing body to prevent resource leak

	return slips.Slips.Advice
}
