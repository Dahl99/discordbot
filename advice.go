package discordbot

import (
	"encoding/json"
	"log"
	"net/http"
)

//Struct used to store advice in json
type slip struct {
	Advice string `json:"advice"`
}

//Struct containing the advice slip
type allSlips struct {
	Slips slip `json:"slip"`
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
		return decodingFailed
	}
	res.Body.Close() // Closing body to prevent resource leak

	return slips.Slips.Advice
}
