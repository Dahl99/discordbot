package chess

import (
	"log"
	"time"

	"github.com/notnil/chess/uci"

	"discordbot/src/models"
)

var eng *uci.Engine

func InitChessAi() {
	// set up engine to use stockfish exe
	var err error
	eng, err = uci.New("stockfish")
	if err != nil {
		log.Println(err)
	}
}

func StopChessAi() {
	eng.Close()
}

func (s *chessSession) getAiMove() string {
	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		log.Println(err)
	}

	cmdPos := uci.CmdPosition{Position: s.game.Position()}
	cmdGo := uci.CmdGo{MoveTime: time.Second * 10}

	if err := eng.Run(cmdPos, cmdGo); err != nil {
		log.Println(err)
	}

	return eng.SearchResults().BestMove.String()
}

func (s *chessSession) isAiTurn(botID string) bool {
	if s.model.GameState == models.TurnWhite && s.isAiPlayerWhite(botID) {
		return true
	} else if s.model.GameState == models.TurnBlack && s.isAiPlayerBlack(botID) {
		return true
	}

	return false
}

func (s *chessSession) isAiPlayerWhite(botID string) bool {
	if s.model.PlayerWhite == botID {
		return true
	}

	return false
}

func (s *chessSession) isAiPlayerBlack(botID string) bool {
	if s.model.PlayerBlack == botID {
		return true
	}

	return false
}

func (s *chessSession) aiMovePiece() {
	aiMove := s.getAiMove()
	err := s.game.MoveStr(aiMove)
	if err != nil {
		log.Println(err)
		return
	}

	s.updateBoardState()
	s.updateGameState()
}
