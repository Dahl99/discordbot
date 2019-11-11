package discordbot

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

//Const containing the root of the url
const scryfallBaseURL string = "https://api.scryfall.com/cards/"

//Const containing string to be sent if scryfall API is unavailable
const scryfallNotAvailable string = "Scryfall API not available at the moment. Please try again later."

//Struct to store info from first http.Get() and autocomplete search
type autoResult struct {
	Data []string `json:"data"`
}

//Sub struct in exactResult struct. It's used to store the imageURIs from scryfall api
type imageURI struct {
	Png string `json:"png"`
}

//Struct used to store data from second http.Get()
type exactResult struct {
	Image imageURI `json:"image_uris"`
}

//getCard() fetches a card based on which card name used in command
func getCard(n []string) string {

	name := replaceSpace(removeOrdMatter(n, 0))       // Replaces the spaces with "_" to avoid url problems
	URL := scryfallBaseURL + "autocomplete?q=" + name // Sets url for autocomplete get request

	res, err := http.Get(URL) // Fetching most probable card using scryfall autocomplete
	if err != nil {           // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return scryfallNotAvailable
	}

	//	Decoding results into autoresult struct object
	var autoresult autoResult
	err = json.NewDecoder(res.Body).Decode(&autoresult)
	if err != nil {
		log.Println(err)
		return decodingFailed
	}
	res.Body.Close() // Closing body to prevent resource leak

	cardName := strings.Split(autoresult.Data[0], " ") // Splitting string on each space
	name = replaceSpace(cardName)                      // Replaceing space with "_" to avoid url problems
	URL = scryfallBaseURL + "named?exact=" + name      // Sets url for exact card get request

	res, err = http.Get(URL) // Fetching exact card
	if err != nil {          // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return scryfallNotAvailable
	}

	// Decoding results into exactResult
	var card exactResult
	err = json.NewDecoder(res.Body).Decode(&card)
	if err != nil {
		log.Println(err)
		return decodingFailed
	}
	res.Body.Close() // Closing body to prevent resource leak

	return card.Image.Png // Returning url to png version of card
}

//	This function replaces spaces in a string slice with "_"
func replaceSpace(s []string) string {
	var retString string //	String to be returned

	i := 0

	if len(s) >= 2 { // Checks if name is more than one word
		for i < len(s) { // Loops through slice and adds index
			retString += s[i]

			if i != len(s)-1 { // If current index isn't last, "_" is appended
				retString += "_"
			}

			i++ // Counts up
		}
	} else { // If name is 1 word, name is set
		retString = s[0]
	}

	return retString // Returns name with "_" instead of spaces or 1 word name
}

//	Removes an index from a string slice while keeping same order
func removeOrdMatter(s []string, ind int) []string {
	return append(s[:ind], s[ind+1:]...)
}
