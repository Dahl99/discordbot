package commands

import (
	"math/rand"
	"strconv"

	"discordbot/src/utils"

	"github.com/bwmarrin/discordgo"
)

func RollDice(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) == 2 { // Checks if die command has correct length
		if _, err := strconv.Atoi(cmd[1]); err == nil { //	Checks if user entered a number
			dieSides, _ := strconv.Atoi(cmd[1]) // Converts die sides to int from ASCII

			if dieSides >= 2 {
				rolled := strconv.Itoa(rand.Intn(dieSides-1) + 1) // Rolls die and returns result as string
				utils.SendChannelMessage(m.ChannelID, m.Author.Mention()+" rolled "+rolled)
			}

		} else {
			utils.SendChannelMessage(m.ChannelID, m.Author.Mention()+" "+cmd[1]+" is not a number!")
		}
	}
}
