package handlers

import (
	"discordbot/src/commands"
	"discordbot/src/config"
	"discordbot/src/consts"
	"discordbot/src/music"
	"discordbot/src/utils"

	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called when the bot receives the "ready" event from Discord.
func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateGameStatus(0, config.Config.Status)
}


// This function will be called every time a new guild is joined.
func GuildCreateHandler(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			s.ChannelMessageSend(channel.ID, config.Config.Online)
			return
		}
	}
}


//MessageCreate will be called everytime a new message is sent in a channel the bot has access to
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // Preventing bot from using own commands
		return
	}

	guildID := utils.SearchGuild(m.ChannelID)
	v := music.VoiceInstances[guildID]
	cmd := strings.Split(m.Content, " ") //	Splitting command into string slice

	switch cmd[0] {
	case config.Config.Prefix + "help":
		utils.SendChannelMessage(m, consts.Help + consts.MusicHelp)
	case config.Config.Prefix + "ping":
		utils.SendChannelMessage(m, "Pong!")
	case config.Config.Prefix + "card":
		commands.PostCard(cmd, m)
	case config.Config.Prefix + "dice":
		commands.RollDice(cmd, m)
	case config.Config.Prefix + "insult":
		commands.PostInsult(m)
	case config.Config.Prefix + "advice":
		commands.PostAdvice(m)
	case config.Config.Prefix + "music":
		music.Music(cmd[1:], v, s, m)
	case config.Config.Prefix + "kanye":
		commands.PostKanyeQuote(m)
	default:
		return
	}
}
