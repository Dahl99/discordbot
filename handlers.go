package discordbot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called when the bot receives the "ready" event from Discord.
func Ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateGameStatus(0, Bot.Status)
}


//MessageCreate will be called everytime a new message is sent in a channel the bot has access to
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // Preventing bot from using own commands
		return
	}

	cmd := strings.Split(m.Content, " ") //	Splitting command into string slice

	switch cmd[0] {
	case Bot.Prefix + "help":
		s.ChannelMessageSend(m.ChannelID, help)
	case Bot.Prefix + "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case Bot.Prefix + "card":
		postCard(cmd, s, m)
	case Bot.Prefix + "dice":
		rollDice(cmd, s, m)
	case Bot.Prefix + "insult":
		postInsult(cmd, s, m)
	case Bot.Prefix + "advice":
		postAdvice(cmd, s, m)
	default:
		return
	}
}


// This function will be called every time a new guild is joined.
func GuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, Bot.Guildjoin)
			return
		}
	}
}