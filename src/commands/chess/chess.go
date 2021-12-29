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

	pngFilepath := strings.TrimRight(filepath, ".svg") + ".png"

	err = exec.Command("inkscape", "-w", "360", "-h", "360", filepath, "-o", pngFilepath).Run()
	if err != nil {
		return ""
	}

	return pngFilepath
}

func (s *chessSession) updateTurnPlayer() {
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
	pngFilepath := s.saveChessBoardToPng()
	if pngFilepath != "" {
		utils.SendChannelFile(channelID, pngFilepath, "board.png")
	}
	exec.Command("rm", pngFilepath).Run()
}

func (s *chessSession) hasWhiteWon() bool {
	if s.game.Outcome() == chess.WhiteWon {
		return true
	}

	return false
}

func (s *chessSession) hasBlackWon() bool {
	if s.game.Outcome() == chess.BlackWon {
		return true
	}

	return false
}

func (s *chessSession) isGameDraw() bool {
	if s.game.Outcome() == chess.Draw {
		return true
	}

	return false
}

func (s *chessSession) isGameOver() bool {
	if s.hasWhiteWon() || s.hasBlackWon() || s.isGameDraw() {
		return true
	}

	return false
}

func (s *chessSession) setWinnerOrDraw() {
	if s.hasWhiteWon() {
		s.model.GameState = models.WonWhite
	} else if s.hasBlackWon() {
		s.model.GameState = models.WonBlack
	} else {
		s.model.GameState = models.Draw
	}
}

func (s *chessSession) endGame(channelID string) {
	s.setWinnerOrDraw()

	message := "**[Chess]** Game is over, "

	if s.hasWhiteWon() {
		message += "<@" + s.model.PlayerWhite + "> has won!"
	} else if s.hasBlackWon() {
		message += "<@" + s.model.PlayerBlack + "> has won!"
	} else {
		message += "it ended as a draw."
	}

	s.model.Update()
	database.DB.Delete(s.model)
	utils.SendChannelMessage(channelID, message)
}
