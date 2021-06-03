package discordbot

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

//Sub struct in exactResult struct. It's used to store the imageURIs from scryfall api
type imageURI struct {
	Png string `json:"png"`
}

type cardFaces struct {
	Image imageURI `json:"image_uris"`
}

type prices struct {
	Usd     string `json:"usd"`
	UsdFoil string `json:"usd_foil"`
}

//Struct used to store data from second http.Get()
type fuzzyResult struct {
	Name   string       `json:"name"`
	Image  imageURI     `json:"image_uris"`
	Prices prices       `json:"prices"`
	Faces  [2]cardFaces `json:"card_faces"`
}

func postCard(cmd []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	// Getting card if card name was entered
	if len(cmd) > 1 {
		s.ChannelMessageSend(m.ChannelID, getCard(cmd))
	}
}

//getCard() fetches a card based on which card name used in command
func getCard(n []string) string {

	name := replaceSpace(n[1:]) // Replaces the spaces with "_" to avoid url problems
	if len(name) < 3 {
		return "Name needs to have 3 or more letters to search"
	}

	URL := scryfallBaseURL + name // Sets url for exact card get request

	res, err := http.Get(URL) // Fetching exact card
	if err != nil {           // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return scryfallNotAvailable
	}

	// Decoding results into exactResult
	var card fuzzyResult
	err = json.NewDecoder(res.Body).Decode(&card)
	if err != nil {
		log.Println(err)
		return decodingFailed
	}

	if card.Image.Png == "" && card.Faces[0].Image.Png == "" && card.Faces[1].Image.Png == "" {
		return "Unable to find requested card, avoid ambigous searches!"
	}

	res.Body.Close() // Closing body to prevent resource leak

	//	Making the returned string
	var result string

	if card.Prices.Usd != "" || card.Prices.UsdFoil != "" {
		result += "\nTCGPlayer price:"
	}

	if card.Prices.Usd != "" {
		result += "\n\tUSD = " + card.Prices.Usd
	}

	if card.Prices.UsdFoil != "" {
		result += "\n\tUSD Foil = " + card.Prices.UsdFoil
	}

	if card.Image.Png != "" {
		result += "\n" + card.Image.Png
	} else {
		result += "\n" + card.Faces[0].Image.Png + "\n" + card.Faces[1].Image.Png
	}

	return result	// Returning url to png version of card
}

//	This function replaces spaces in a string slice with "_"
func replaceSpace(s []string) string {
	var result string //	String to be returned

	i := 0

	if len(s) > 1 { // Checks if name is more than one word
		for i < len(s) { // Loops through slice and adds index
			result += s[i]

			if i != len(s)-1 { // If current index isn't last, "_" is appended
				result += "_"
			}

			i++ // Counts up
		}
	} else { // If name is 1 word, name is set
		result = s[0]
	}

	return result // Returns name with "_" instead of spaces or 1 word name
}
