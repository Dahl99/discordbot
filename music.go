package discordbot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func music(cmd []string, s *discordgo.Session, m *discordgo.MessageCreate) {

	if len(cmd) < 2 {
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

func playMusic(n []string, s *discordgo.Session, m *discordgo.MessageCreate) {

	name := replaceSpace(n)
	url := youtubeEndpoint + Bot.Ytkey + "&q=" + name

	result := strings.Split(ytSearch(url), "|")

	if result[0] == ytSearchFailed || result[0] == decodingFailed {
		s.ChannelMessageSend(m.ChannelID, result[0])
	}

	s.ChannelMessageSend(m.ChannelID, result[0] + "\t" + result[1])
}

func skipMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Skip music!")
}
