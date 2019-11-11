package discordbot

import (
	"encoding/json"
	"log"
	"net/http"
)

//insultURL contains the url for the API generating insults
const insultURL string = "https://evilinsult.com/generate_insult.php?lang=en&type=json"

//String to be sent if Evil Insult API isn't available
const evilInsultNotAvailable string = "Evil Insult API not available at the moment. Please try again later."

// Struct to store fetched data from Evil Insult API
type insult struct {
	Insult string `json:"insult"`
}

func getInsult() string {
	res, err := http.Get(insultURL) // Fetching most probable card using scryfall autocomplete
	if err != nil {                 // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return evilInsultNotAvailable
	}

	var insultObj insult

	err = json.NewDecoder(res.Body).Decode(&insultObj) // Decoding data into struct object
	if err != nil {
		log.Println(err)
		return decodingFailed
	}

	res.Body.Close() // Closing body to prevent resource leak

	return insultObj.Insult
}
