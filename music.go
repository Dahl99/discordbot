package discordbot

import (
	"github.com/bwmarrin/discordgo"
)

func music(cmd []string, s *discordgo.Session, m *discordgo.MessageCreate) {

	if len(cmd) < 3 {
		return
	}

	switch(cmd[1]) {
	case "play":
		playMusic(s, m)
	case "skip":
		skipMusic(s, m)
	default:
		return
	}
}

func playMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Play music!")
}

func skipMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Skip music!")
}
