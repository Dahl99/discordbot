package discordbot

// searchGuild search the guild ID
func searchGuild(textChannelID string) (guildID string) {
	channel, _ := dg.Channel(textChannelID)
	guildID = channel.GuildID
	return guildID
}

// searchVoiceChannel search the voice channel id into from guild.
func searchVoiceChannel(user string) (voiceChannelID string) {
	for _, g := range dg.State.Guilds {
		for _, v := range g.VoiceStates {
			if v.UserID == user {
				return v.ChannelID
			}
		}
	}
	return ""
}
