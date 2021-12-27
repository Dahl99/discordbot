package chess

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/notnil/chess"

	"discordbot/src/database"
	"discordbot/src/models"
	"discordbot/src/utils"
)

type challenge struct {
	guildID    string
	challenger string
	opponent   string
}

var challenges []*challenge

func createNewGame(index int, channelID string) {
	challenge := challenges[index]

	var playerWhite string
	var playerBlack string
	roll := rand.Intn(1)
	if roll == 0 {
		playerWhite = challenge.challenger
		playerBlack = challenge.opponent
	} else {
		playerBlack = challenge.challenger
		playerWhite = challenge.opponent
	}

	chessGame := chess.NewGame()

	chessGameModel := &models.ChessGame{
		GuildID:     challenge.guildID,
		PlayerWhite: playerWhite,
		PlayerBlack: playerBlack,
		BoardState:  chessGame.String(),
		GameState:   models.TurnWhite,
	}

	database.DB.Create(chessGameModel)

	utils.SendChannelMessage(channelID, "**[Chess]** Game has been started, <@"+playerWhite+"> make the first move!")
}

func movePiece(m *discordgo.MessageCreate, move string) {

}
