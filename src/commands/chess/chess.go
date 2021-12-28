package chess

import (
	"log"
	"math/rand"
	"os"
	"os/exec"
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
	session.model = &models.ChessGame{
		GuildID:     challenge.guildID,
		PlayerWhite: playerWhite,
		PlayerBlack: playerBlack,
		BoardState:  session.game.String(),
		GameState:   models.TurnWhite,
	}

	database.DB.Create(session.model)

	if session.isAiPlayerWhite(botID) {
		utils.SendChannelMessage(channelID, "**[Chess]** Game has started, <@"+botID+"> is moving the first piece!")
		session.aiMovePiece()
		session.model.Update()
	} else {
		utils.SendChannelMessage(channelID, "**[Chess]** Game has been started, <@"+playerWhite+"> move the first piece!")
	}

	session.sendChannelChessBoard(channelID)
}

func (s *chessSession) saveChessBoardToPng() string {
	filepath := s.model.GuildID + "-" + strconv.FormatInt(s.model.CreatedAt, 10)
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer f.Close()

	// create board position
	fenStr := s.game.FEN()
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

	err = exec.Command("inkscape", "-w", "360", "-h", "360", filepath, "-o", pngName).Run()
	if err != nil {
		return ""
	}

	return pngName
}

func (s *chessSession) updateGameState() {
	if s.model.GameState == models.TurnWhite {
		s.model.GameState = models.TurnBlack
	} else {
		s.model.GameState = models.TurnWhite
	}
}

func (s *chessSession) updateBoardState() {
	s.model.BoardState = s.game.String()
}

func (s *chessSession) sendChannelChessBoard(channelID string) {
	png := s.saveChessBoardToPng()
	if png != "" {
		utils.SendChannelFile(channelID, png, "board.png")
	}
	exec.Command("rm", png).Run()
}
