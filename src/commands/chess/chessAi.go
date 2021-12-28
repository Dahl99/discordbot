package chess

import (
	"log"
	"time"

	"github.com/notnil/chess/uci"
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

func aiMovePiece(session *chessSession) *chessSession {
	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		log.Println(err)
	}

	cmdPos := uci.CmdPosition{Position: session.game.Position()}
	cmdGo := uci.CmdGo{MoveTime: time.Second / 150}

	if err := eng.Run(cmdPos, cmdGo); err != nil {
		log.Println(err)
	}

	move := eng.SearchResults().BestMove
	session.game.MoveStr(move.String())

	session.model.BoardState = session.game.String()
	session.model.Update()

	return session
}
