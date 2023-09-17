package handlers

import (
	"github.com/Dahl99/discordbot/internal/commands/music"

	"github.com/Dahl99/discordbot/internal/commands"
	"github.com/Dahl99/discordbot/internal/commands/chess"
	"github.com/Dahl99/discordbot/internal/config"
	"github.com/Dahl99/discordbot/internal/discord"

	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// help is a constant for info provided in help command.
const help string = "```Current commands are:\n\tping\n\tcard <card name>\n\tdice <die sides>\n\tinsult\n\tadvice\n\tkanye"

// musicHelp is a constant for info provided about music functionality in help command.
const musicHelp string = "\n\nMusic commands:\n\tplay <youtube url/query>\n\tleave\n\tskip\n\tstop"

const chessHelp string = "\n\nChess commands (prefix chess command):\n\tchallenge @opponent\n\taccept\n\tdecline\n\tmove <algebraic notation>\n\tresign```"

// ReadyHandler will be called when the bot receives the "ready" event from Discord.
func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	err := s.UpdateGameStatus(0, config.GetBotStatus())
	if err != nil {
		slog.Warn("failed to update game status", "error", err)
	}
}

// GuildCreateHandler will be called every time a new guild is joined.
func GuildCreateHandler(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, err := s.ChannelMessageSend(channel.ID, config.GetBotGuildJoinMessage())
			if err != nil {
				slog.Warn("failed to send guild create handler message", "error", err)
			}

			return
		}
	}
}

// MessageCreateHandler will be called everytime a new message is sent in a channel the bot has access to.
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // Preventing bot from using own commands
		return
	}

	slog.Info("processing command", "command", m.Content)

	prefix := config.GetBotPrefix()
	guildID := discord.SearchGuildByChannelID(m.ChannelID)
	v := music.VoiceInstances[guildID]
	cmd := strings.Split(m.Content, " ") //	Splitting command into string slice

	switch cmd[0] {
	case prefix + "help":
		discord.SendChannelMessage(m.ChannelID, help+musicHelp+chessHelp)
	case prefix + "ping":
		discord.SendChannelMessage(m.ChannelID, "Pong!")
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
	case prefix + "play":
		music.PlayMusic(cmd[1:], v, s, m)
	case prefix + "leave":
		music.LeaveVoice(v, m)
	case prefix + "skip":
		music.SkipMusic(v, m)
	case prefix + "stop":
		music.StopMusic(v, m)
	case prefix + "chess":
		chess.Menu(cmd[1:], s, m)
	default:
		return
	}
}
