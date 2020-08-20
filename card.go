package discordbot

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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
	if len(cmd) == 1 { // Checks if card name is missing
		log.Println("Missing card name!")
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" missing card name!")
	} else {
		s.ChannelMessageSend(m.ChannelID, getCard(cmd))
	}
}

//getCard() fetches a card based on which card name used in command
func getCard(n []string) string {

	name := replaceSpace(removeOrdMatter(n)) // Replaces the spaces with "_" to avoid url problems
	if len(name) <= 2 {
		return "Name needs to have 3 or more letters to search"
	}

	URL := scryfallBaseURL + name // Sets url for exact card get request

	res, err := http.Get(URL) // Fetching exact card
	if err != nil {           // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return scryfallNotAvailable
	}

	time.Sleep(200 * time.Millisecond) // Sleeping for 0,2 seconds to prevent spam

	// Decoding results into exactResult
	var card fuzzyResult
	err = json.NewDecoder(res.Body).Decode(&card)
	if err != nil {
		log.Println(err)
		return decodingFailed
	}

	if card.Image.Png == "" && card.Faces[0].Image.Png == "" && card.Faces[1].Image.Png == "" {
		return "Unable to find requested card"
	}

	res.Body.Close() // Closing body to prevent resource leak

	//	Making the returned string
	var retString string

	if card.Prices.Usd != "" || card.Prices.UsdFoil != "" {
		retString += "\nTCGPlayer price:"
	}

	if card.Prices.Usd != "" {
		retString += "\n\tUSD = " + card.Prices.Usd
	}

	if card.Prices.UsdFoil != "" {
		retString += "\n\tUSD Foil = " + card.Prices.UsdFoil
	}

	if card.Image.Png != "" {
		retString += "\n" + card.Image.Png
	} else {
		retString += "\n" + card.Faces[0].Image.Png + "\n" + card.Faces[1].Image.Png
	}

	return retString // Returning url to png version of card
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
func removeOrdMatter(s []string) []string {
	return append(s[:0], s[1:]...)
}
