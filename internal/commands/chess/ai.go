package chess

import (
	"github.com/Dahl99/discord-bot/internal/models"
	"log/slog"
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

var eng *uci.Engine

func InitAi() {
	// set up engine to use stockfish exe
	var err error
	eng, err = uci.New("stockfish")
	if err != nil {
		slog.Warn("failed to create chess engine", "error", err)
	}
}

func StopChessAi() {
	eng.Close()
}

func (s *chessSession) getAiMove() *chess.Move {
	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		slog.Warn("failed to run chess command", "error", err)
	}

	cmdPos := uci.CmdPosition{Position: s.game.Position()}
	cmdGo := uci.CmdGo{MoveTime: time.Second * 10}

	if err := eng.Run(cmdPos, cmdGo); err != nil {
		slog.Warn("failed to run chess command", "error", err)
	}

	return eng.SearchResults().BestMove
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
	if err := s.game.Move(aiMove); err != nil {
		slog.Warn("failed to do move for chess ai", "error", err)
		return
	}
}
