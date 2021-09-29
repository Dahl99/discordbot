package utils

import (
	"discordbot/src/config"

	"github.com/bwmarrin/discordgo"
)

// SearchGuild search the guild ID
func SearchGuild(textChannelID string) (guildID string) {
	channel, _ := config.Dg.Channel(textChannelID)
	guildID = channel.GuildID
	return guildID
}

// SearchVoiceChannel search the voice channel id into from guild.
func SearchVoiceChannel(user string) (voiceChannelID string) {
	for _, g := range config.Dg.State.Guilds {
		for _, v := range g.VoiceStates {
			if v.UserID == user {
				return v.ChannelID
			}
		}
	}
	return ""
}

// SendChannelMessage sends a channel message to channel with channel id equal to m.ChannelID
func SendChannelMessage(m *discordgo.MessageCreate, message string) {
	config.Dg.ChannelMessageSend(m.ChannelID, message)
}
