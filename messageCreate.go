package discordbot

import (
	"log"
	"strconv"
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
		if len(cmd) == 1 { // Checks if card name is missing
			log.Println("Missing card name!")
			s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" missing card name!")
		} else {
			s.ChannelMessageSend(m.ChannelID, getCard(cmd))
		}
	case prefix + "dice":
		if len(cmd) == 2 && cmd[1] >= "2" { // Checks if die command has correct length
			if _, err := strconv.Atoi(cmd[1]); err == nil { //	Checks if user entered a number
				s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" rolled "+diceRoll(cmd))
			} else {
				s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" "+cmd[1]+" is not a number!")
			}
		}
	case prefix + "insult":
		if len(cmd) == 1 { // Checks if insult command was used properly
			s.ChannelMessageSend(m.ChannelID, getInsult())
		}
	case prefix + "advice":
		if len(cmd) == 1 { // Checks if advice command was used properly
			s.ChannelMessageSend(m.ChannelID, getAdvice())
		}
	default:
		return
	}
}
