package commands

import (
	"discordbot/src/bot"
	"discordbot/src/consts"
	"encoding/json"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

//Struct used to store advice in json
type slip struct {
	Advice string `json:"advice"`
}

//Struct containing the advice slip
type allSlips struct {
	Slips slip `json:"slip"`
}

func PostAdvice(m *discordgo.MessageCreate) {
	bot.Dg.ChannelMessageSend(m.ChannelID, getAdvice())
}

func getAdvice() string {
	res, err := http.Get(consts.AdviceSlipURL) // Fetching an advice
	if err != nil {                     // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return consts.AdviceslipNotAvailable
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
