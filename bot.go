package discordbot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type bot struct {
	Token string
	Status string
}


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

//MessageCreate will be called everytime a new message is sent in a channel the bot has access to
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // Preventing bot from using own commands
		return
	}

	cmd := strings.Split(m.Content, " ") //	Splitting command into string slice

	switch cmd[0] {
	case prefix + "help":
		s.ChannelMessageSend(m.ChannelID, help)
	case prefix + "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case prefix + "card":
		postCard(cmd, s, m)
	case prefix + "dice":
		rollDice(cmd, s, m)
	case prefix + "insult":
		postInsult(cmd, s, m)
	case prefix + "advice":
		postAdvice(cmd, s, m)
	default:
		return
	}
}
