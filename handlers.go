package discordbot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called when the bot receives the "ready" event from Discord.
func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateGameStatus(0, conf.Status)
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
	case conf.Prefix + "help":
		s.ChannelMessageSend(m.ChannelID, help)
	case conf.Prefix + "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case conf.Prefix + "card":
		postCard(cmd, s, m)
	case conf.Prefix + "dice":
		rollDice(cmd, s, m)
	case conf.Prefix + "insult":
		postInsult(cmd, s, m)
	case conf.Prefix + "advice":
		postAdvice(cmd, s, m)
	case conf.Prefix + "music":
		music(cmd[1:], v, s, m)
	default:
		return
	}
}


// This function will be called every time a new guild is joined.
func GuildCreateHandler(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, conf.Online)
			return
		}
	}
}

// searchGuild search the guild ID
func searchGuild(textChannelID string) (guildID string) {
	channel, _ := dg.Channel(textChannelID)
	guildID = channel.GuildID
	return guildID
}
