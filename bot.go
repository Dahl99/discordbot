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
	Guildjoin string
}

// Global struct object containing bot config
var Bot bot

// readJsonBotData reads the data the bot needs from the provided JSON file
func ReadJsonBotData() bot {
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

