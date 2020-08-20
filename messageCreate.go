package discordbot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

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
	case prefix + "play":
		if len(cmd) == 2 {
			playMusic(cmd[1])
		}
	default:
		return
	}
}
