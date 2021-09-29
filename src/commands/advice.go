package commands

import (
	"encoding/json"
	"log"
	"net/http"

	"discordbot/src/consts"
	"discordbot/src/utils"

	"github.com/bwmarrin/discordgo"
)

// adviceSlipURL contains url to adviceslip API
const adviceSlipURL string = "https://api.adviceslip.com/advice"

// adviceslipNotAvailable contains string to be sent if adviceslip API is unavailable
const adviceslipNotAvailable string = "Adviceslip API not available at the moment."

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
	if err != nil {                     // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return adviceslipNotAvailable
	}

	//	Decoding results into autoresult struct object
	var slips allSlips
	err = json.NewDecoder(res.Body).Decode(&slips)
	if err != nil {
		log.Println(err)
		return consts.DecodingFailed
	}
	res.Body.Close() // Closing body to prevent resource leak

	return slips.Slips.Advice
}
