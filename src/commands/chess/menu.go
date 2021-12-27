package chess

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"discordbot/src/utils"
)

func Menu(cmd []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	switch cmd[0] {
	case "challenge":
		challengePlayer(s, m, cmd[1])
	case "accept":
		accept(m.ID, m.GuildID, m.ChannelID)
	case "mode":
		movePiece(m, cmd[1])
	default:
		return
	}
}

func challengePlayer(s *discordgo.Session, challenger *discordgo.MessageCreate, opponent string) {
	utils.SendChannelMessage(challenger.ChannelID, "**[Chess]** "+opponent+
		" you have been challenged to a game by <@"+challenger.Author.ID+"> do you accept?")

	opponentID := opponent
	opponentID = strings.TrimLeft(opponentID, "<@!")
	opponentID = strings.TrimRight(opponentID, ">")

	challenge := &challenge{
		guildID:    challenger.GuildID,
		challenger: challenger.Author.ID,
		opponent:   opponentID,
	}

	challenges = append(challenges, challenge)

	if opponentID == s.State.User.ID {
		accept(s.State.User.ID, challenger.GuildID, challenger.ChannelID)
	}
}

func accept(userID string, guildID string, channelID string) {
	for index, challenge := range challenges {
		if userID == challenge.opponent && guildID == challenge.guildID {
			createNewGame(index, channelID)
			utils.SendChannelMessage(channelID, "**[Chess]** Challenge accepted, starting new game")
		}
	}
}
