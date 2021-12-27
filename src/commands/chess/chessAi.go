package chess

import (
	"log"
	"time"

	"github.com/notnil/chess"
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

func aiMove(game *chess.Game) string {
	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		log.Println(err)
	}

	cmdPos := uci.CmdPosition{Position: game.Position()}
	cmdGo := uci.CmdGo{MoveTime: time.Second / 100}

	if err := eng.Run(cmdPos, cmdGo); err != nil {
		log.Println(err)
	}

	move := eng.SearchResults().BestMove

	return move.String()
}
