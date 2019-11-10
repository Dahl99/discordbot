package discordbot

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "+"

//MessageCreate will be called everytime a new message is sent in a channel the bot has access to
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	cmd := strings.Split(m.Content, " ")

	switch cmd[0] {
	case prefix + "Ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case prefix + "Card":
		if len(cmd) == 1 { // Checks if card name is missing
			log.Println("Missing card name!")
			s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" missing card name!")
		} else {
			s.ChannelMessageSend(m.ChannelID, getCard(cmd))
		}
	default:
		return
	}
}
