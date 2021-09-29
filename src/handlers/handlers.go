package handlers

import (
	"discordbot/src/commands"
	"discordbot/src/config"
	"discordbot/src/consts"
	"discordbot/src/context"
	"discordbot/src/music"
	"discordbot/src/utils"

	"strings"

	"github.com/bwmarrin/discordgo"
)

func AddHandlers() {
	// Register handlers as callbacks for the events.
	context.Dg.AddHandler(ReadyHandler)
	// context.Dg.AddHandler(GuildCreateHandler)
	context.Dg.AddHandler(MessageCreateHandler)
}

// ReadyHandler will be called when the bot receives the "ready" event from Discord.
func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateGameStatus(0, config.GetStatusText())
}

// GuildCreateHandler will be called every time a new guild is joined.
func GuildCreateHandler(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			s.ChannelMessageSend(channel.ID, config.GetOnlineText())
			return
		}
	}
}

// MessageCreateHandler will be called everytime a new message is sent in a channel the bot has access to
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // Preventing bot from using own commands
		return
	}

	prefix := config.GetPrefix()
	guildID := utils.SearchGuild(m.ChannelID)
	v := music.VoiceInstances[guildID]
	cmd := strings.Split(m.Content, " ") //	Splitting command into string slice

	switch cmd[0] {
	case prefix + "help":
		utils.SendChannelMessage(m.ChannelID, consts.Help+consts.MusicHelp)
	case prefix + "ping":
		utils.SendChannelMessage(m.ChannelID, "Pong!")
	case prefix + "card":
		commands.PostCard(cmd, m)
	case prefix + "dice":
		commands.RollDice(cmd, m)
	case prefix + "insult":
		commands.PostInsult(m)
	case prefix + "advice":
		commands.PostAdvice(m)
	case prefix + "kanye":
		commands.PostKanyeQuote(m)
	case prefix + "join":
		music.JoinVoice(v, s, m)
	case prefix + "leave":
		music.LeaveVoice(v, m)
	case prefix + "play":
		music.PlayMusic(cmd[1:], v, m)
	case prefix + "skip":
		music.SkipMusic(v, m)
	case prefix + "stop":
		music.StopMusic(v, m)
	default:
		return
	}
}
