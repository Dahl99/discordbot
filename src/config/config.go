package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
)
type configuration struct {
	Token string
	Prefix string
	Status string
	Online string
	Ytkey string
}

var Dg *discordgo.Session

// Global struct object containing bot config
var Config *configuration

// readconfigig reads the data the bot needs from the provided JSON file
func ReadConfiguration() {
	res, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(res, &Config)
	if err != nil {
		log.Println(err)
	}
}