package discordbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

type bot struct {
	Token string
	Prefix string
	Status string
	Online string
}

// lock used to ensure Bot object is a singleton
var lock = &sync.Mutex{}

// Global struct object containing bot config
var Bot *bot

// Gets the instance of the bot singleton object
func GetInstance() *bot {
	if Bot == nil {
		lock.Lock()
		defer lock.Unlock()

		if Bot == nil {
			fmt.Println("Creating single bot instance now")
			Bot = &bot{}
			Bot = readConfig()
		} else {
			fmt.Println("Single instance already created!-1")
		}
	} else {
		fmt.Println("Single instance already created!-2")
	}

	return Bot
}

// readConfig reads the data the bot needs from the provided JSON file
func readConfig() *bot {
	res, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Println(err)
	}

	var temp bot

	err = json.Unmarshal(res, &temp)
	if err != nil {
		log.Println(err)
	}

	return &temp
}

