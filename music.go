package discordbot

import (
	"github.com/bwmarrin/discordgo"
)

func music(cmd []string, s *discordgo.Session, m *discordgo.MessageCreate) {

	if len(cmd) < 3 {
		return
	}

	switch(cmd[0]) {
	case "play":
		playMusic(cmd[1:], s, m)
	case "skip":
		skipMusic(s, m)
	default:
		return
	}
}

func playMusic(name []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Play music!")
}

func skipMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Skip music!")
}
