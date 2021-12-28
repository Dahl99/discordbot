package chess

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/notnil/chess"

	"discordbot/src/database"
	"discordbot/src/models"
	"discordbot/src/utils"
)

func Menu(cmd []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	switch cmd[0] {
	case "challenge":
		challengePlayer(s, m, cmd[1])
	case "accept":
		accept(m.ID, m.GuildID, m.ChannelID, s.State.User.ID)
	case "move":
		movePiece(m, cmd[1], s.State.User.ID)
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
		accept(s.State.User.ID, challenger.GuildID, challenger.ChannelID, s.State.User.ID)
	}
}

func accept(userID string, guildID string, channelID string, botID string) {
	for index, challenge := range challenges {
		if userID == challenge.opponent && guildID == challenge.guildID {
			createNewGame(index, channelID, botID)
		}
	}
}

func movePiece(m *discordgo.MessageCreate, move string, botID string) {
	var session chessSession

	database.DB.Raw(
		"SELECT * "+
			"FROM chess_games "+
			"WHERE guild_id = ? && (player_white = ? || player_black = ?)",
		m.GuildID, m.Author.ID, m.Author.ID).Scan(&session.model)

	if !(session.model.GameState == models.TurnBlack && session.model.PlayerBlack == m.Author.ID) &&
		!(session.model.GameState == models.TurnWhite && session.model.PlayerWhite == m.Author.ID) {
		return
	}

	pgnReader := strings.NewReader(session.model.BoardState)
	pgn, err := chess.PGN(pgnReader)
	if err != nil {
		log.Println("ERR: PGN creation failed")
	}

	session.game = chess.NewGame(pgn)
	err = session.game.MoveStr(move)
	if err != nil {
		log.Println(err)
	}

	if session.model.GameState == models.TurnWhite {
		session.model.GameState = models.TurnBlack
	} else {
		session.model.GameState = models.TurnWhite
	}

	if session.model.GameState == models.TurnBlack && session.model.PlayerBlack == botID {
		aiMove := getAiMove(&session)
		session.game.MoveStr(aiMove)
		session.model.GameState = models.TurnWhite
	}

	if session.model.GameState == models.TurnWhite && session.model.PlayerWhite == botID {
		aiMove := getAiMove(&session)
		session.game.MoveStr(aiMove)
		session.model.GameState = models.TurnBlack
	}

	session.model.BoardState = session.game.String()
	session.model.Update()

	filepath := saveChessBoardToPng(&session)
	if filepath != "" {
		utils.SendChannelFile(m.ChannelID, filepath, "board.png")
	}
}
