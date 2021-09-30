package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type configuration struct {
	Token  string
	Prefix string
	Status string
	Online string
	Ytkey  string
}

// Config is a global struct object containing bot config
var config *configuration

func GetDiscordToken() string {
	return config.Token
}

func GetPrefix() string {
	return config.Prefix
}

func GetStatusText() string {
	return config.Status
}

func GetOnlineText() string {
	return config.Online
}

func GetYtKey() string {
	return config.Ytkey
}

// Load reads the data the bot needs from the provided JSON file
func Load() {
	res, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(res, &config)
	if err != nil {
		log.Fatalln(err)
	}
}
