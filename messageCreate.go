package discordbot

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//Const containing the prefix needed to use bot commands
const prefix string = "+"

//Const containing string to be sent if decoding fails
const decodingFailed string = "Something wrong happened when decoding data"

//MessageCreate will be called everytime a new message is sent in a channel the bot has access to
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // Preventing bot from using own commands
		return
	}

	cmd := strings.Split(m.Content, " ") //	Splitting command into string slice

	switch cmd[0] {
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
		if len(cmd) != 2 { // Checks if die command was used properly
			log.Println("Dice command used wrongly!")
			s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" dice command used wrongly!")
		} else {
			s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" rolled "+diceRoll(cmd))
		}
	case prefix + "insult":
		if len(cmd) != 1 { // Checks if die command was used properly
			log.Println("Insult command used wrongly!")
			s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" insult command used wrongly!")
		} else {
			s.ChannelMessageSend(m.ChannelID, getInsult())
		}
	default:
		return
	}
}
