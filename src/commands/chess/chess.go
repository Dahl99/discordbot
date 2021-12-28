package chess

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/notnil/chess"
	"github.com/notnil/chess/image"

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

type chessSession struct {
	model *models.ChessGame
	game  *chess.Game
}

func createNewGame(index int, channelID string, botID string) {
	challenge := challenges[index]

	var playerWhite string
	var playerBlack string
	if rand.Intn(1) == 0 {
		playerWhite = challenge.challenger
		playerBlack = challenge.opponent
	} else {
		playerWhite = challenge.opponent
		playerBlack = challenge.challenger
	}

	var session chessSession
	session.game = chess.NewGame()
	log.Println("Creating new game: " + session.game.String())

	session.model = &models.ChessGame{
		GuildID:     challenge.guildID,
		PlayerWhite: playerWhite,
		PlayerBlack: playerBlack,
		BoardState:  session.game.String(),
		GameState:   models.TurnWhite,
	}

	database.DB.Create(session.model)

	if playerWhite == botID {
		utils.SendChannelMessage(channelID, "**[Chess]** Game has started, <@"+botID+"> is moving the first piece!")
		aiMove := getAiMove(&session)
		session.game.MoveStr(aiMove)
		session.model.BoardState = session.game.String()
		session.model.GameState = models.TurnBlack
		session.model.Update()

		filepath := saveChessBoardToPng(&session)
		if filepath != "" {
			utils.SendChannelFile(channelID, filepath, "index.png")
		}

	} else {
		utils.SendChannelMessage(channelID, "**[Chess]** Game has been started, <@"+playerWhite+"> move the first piece!")
		filepath := saveChessBoardToPng(&session)
		if filepath != "" {
			utils.SendChannelFile(channelID, filepath, "index.png")
		}
	}
}

func saveChessBoardToPng(session *chessSession) string {
	filepath := session.model.GuildID + "-" + strconv.FormatInt(session.model.CreatedAt, 10)
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer f.Close()

	// create board position
	fenStr := session.game.FEN()
	pos := &chess.Position{}
	if err := pos.UnmarshalText([]byte(fenStr)); err != nil {
		log.Fatal(err)
		return ""
	}

	if err := image.SVG(f, pos.Board()); err != nil {
		log.Fatal(err)
		return ""
	}

	pngName := strings.TrimRight(filepath, ".svg") + ".png"
	utils.SVGtoPNG(filepath, pngName)

	return pngName
}
