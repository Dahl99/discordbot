package discordbot

import (
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
	default:
		return
	}
}
