package handlers

import (
	"discordbot/src/bot"
	"discordbot/src/commands"
	"discordbot/src/consts"
	"discordbot/src/music"
	"discordbot/src/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called when the bot receives the "ready" event from Discord.
func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateGameStatus(0, bot.Config.Status)
}


// This function will be called every time a new guild is joined.
func GuildCreateHandler(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			s.ChannelMessageSend(channel.ID, bot.Config.Online)
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
	case bot.Config.Prefix + "help":
		bot.Dg.ChannelMessageSend(m.ChannelID, consts.Help + consts.MusicHelp)
	case bot.Config.Prefix + "ping":
		bot.Dg.ChannelMessageSend(m.ChannelID, "Pong!")
	case bot.Config.Prefix + "card":
		commands.PostCard(cmd, m)
	case bot.Config.Prefix + "dice":
		commands.RollDice(cmd, m)
	case bot.Config.Prefix + "insult":
		commands.PostInsult(m)
	case bot.Config.Prefix + "advice":
		commands.PostAdvice(m)
	case bot.Config.Prefix + "music":
		music.Music(cmd[1:], v, s, m)
	case bot.Config.Prefix + "kanye":
		commands.PostKanyeQuote(m)
	default:
		return
	}
}
