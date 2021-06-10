package discordbot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called when the bot receives the "ready" event from Discord.
func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateGameStatus(0, config.Status)
}


// This function will be called every time a new guild is joined.
func GuildCreateHandler(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			s.ChannelMessageSend(channel.ID, config.Online)
			return
		}
	}
}


//MessageCreate will be called everytime a new message is sent in a channel the bot has access to
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // Preventing bot from using own commands
		return
	}

	guildID := searchGuild(m.ChannelID)
	v := voiceInstances[guildID]
	cmd := strings.Split(m.Content, " ") //	Splitting command into string slice

	switch cmd[0] {
	case config.Prefix + "help":
		dg.ChannelMessageSend(m.ChannelID, help + musicHelp)
	case config.Prefix + "ping":
		dg.ChannelMessageSend(m.ChannelID, "Pong!")
	case config.Prefix + "card":
		postCard(cmd, m)
	case config.Prefix + "dice":
		rollDice(cmd, m)
	case config.Prefix + "insult":
		postInsult(m)
	case config.Prefix + "advice":
		postAdvice(m)
	case config.Prefix + "music":
		music(cmd[1:], v, s, m)
	case config.Prefix + "kanye":
		postKanyeQuote(m)
	default:
		return
	}
}
