package commands

import (
	"encoding/json"
	"github.com/Dahl99/discord-bot/internal/discord"
	"log/slog"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// scryfallBaseUrl contains the root of the url
const scryfallBaseUrl string = "https://api.scryfall.com/cards/named?fuzzy="

// ScryfallNotAvailable contains string to be sent if scryfall API is unavailable
const ScryfallNotAvailable string = "Scryfall API not available at the moment."

// Sub struct in exactResult struct. It's used to store the imageURIs from scryfall api
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

// Struct used to store data from second http.Get()
type fuzzyResult struct {
	Name   string       `json:"name"`
	Image  imageURI     `json:"image_uris"`
	Prices prices       `json:"prices"`
	Faces  [2]cardFaces `json:"card_faces"`
}

func PostCard(cmd []string, m *discordgo.MessageCreate) {
	// Getting card if card name was entered
	if len(cmd) > 1 {
		discord.SendChannelMessage(m.ChannelID, getCard(cmd))
	}
}

// getCard() fetches a card based on which card name used in command
func getCard(n []string) string {

	name := strings.Join(n[1:], "_")

	if len(name) < 3 {
		return "Name needs to have 3 or more letters to search"
	}

	URL := scryfallBaseUrl + name // Sets url for exact card get request

	res, err := http.Get(URL) // Fetching exact card
	if err != nil {
		slog.Warn("failed to get card from Scryfall", "error", err)
		return ScryfallNotAvailable
	}

	// Decoding fuzzyResult from get request
	var card fuzzyResult
	err = json.NewDecoder(res.Body).Decode(&card)
	if err != nil {
		slog.Warn("failed to decode response from Scryfall", "error", err)
		return "ERR: decoding data failed"
	}

	res.Body.Close()

	if card.Image.Png == "" && card.Faces[0].Image.Png == "" && card.Faces[1].Image.Png == "" {
		return "Unable to find requested card, avoid ambiguous searches!"
	}

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

	return result // Returning url to png version of card
}
