package discordbot

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type bot struct {
	Token string
	Prefix string
	Status string
	Online string
	Ytkey string
}

// Global struct object containing bot config
var Bot bot

// ReadConfig reads the data the bot needs from the provided JSON file
func ReadConfig() bot {
	res, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Println(err)
	}

	var bot bot

	err = json.Unmarshal(res, &bot)
	if err != nil {
		log.Println(err)
	}

	return bot
}

