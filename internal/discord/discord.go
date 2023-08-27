package discord

import (
	"github.com/Dahl99/DiscordBot/internal/context"
	"log/slog"
	"os"
)

// SearchGuild search the guild ID
func SearchGuild(textChannelID string) (guildID string) {
	channel, _ := context.Dg.Channel(textChannelID)
	guildID = channel.GuildID
	return guildID
}

// SearchVoiceChannel search the voice channel id into from guild.
func SearchVoiceChannel(user string) (voiceChannelID string) {
	for _, g := range context.Dg.State.Guilds {
		for _, v := range g.VoiceStates {
			if v.UserID == user {
				return v.ChannelID
			}
		}
	}
	return ""
}

// SendChannelMessage sends a channel message to channel with channel id equal to m.ChannelID
func SendChannelMessage(channelID string, message string) {
	_, err := context.Dg.ChannelMessageSend(channelID, message)
	if err != nil {
		slog.Warn("failed to send message to channel", "channelId", channelID, "message", message, "error", err)
	}
}

func SendChannelFile(channelID string, filepath string, name string) {
	reader, err := os.Open(filepath)
	if err != nil {
		slog.Warn("failed to open file", "filepath", filepath, "error", err)
		return
	}

	_, err = context.Dg.ChannelFileSend(channelID, name, reader)
	if err != nil {
		slog.Warn("failed to send file to channel", "channelId", channelID, "filepath", filepath, "error", err)
	}
}
