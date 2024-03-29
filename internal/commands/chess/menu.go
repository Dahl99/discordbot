package chess

import (
	"log/slog"
	"strings"

	"github.com/Dahl99/discordbot/internal/database"
	"github.com/Dahl99/discordbot/internal/discord"
	"github.com/Dahl99/discordbot/internal/models"

	"github.com/bwmarrin/discordgo"
	"github.com/notnil/chess"
)

func Menu(cmd []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	switch cmd[0] {
	case "challenge":
		challengePlayer(s, m, cmd[1])
	case "accept":
		accept(m.Author.ID, m.GuildID, m.ChannelID, s.State.User.ID)
	case "decline":
		decline(m)
	case "move":
		movePiece(m, cmd[1], s.State.User.ID)
	case "resign":
		resign(m)
	default:
		return
	}
}

func challengePlayer(s *discordgo.Session, challenger *discordgo.MessageCreate, opponent string) {
	var count int
	database.DB.Raw(
		"SELECT IFNULL((SELECT COUNT(id) FROM chess_games WHERE player_white = ? OR player_black = ?), 0) "+
			"FROM chess_games WHERE deleted_at IS NULL",
		challenger.Author.ID, challenger.Author.ID).Scan(&count)

	if count > 0 {
		discord.SendChannelMessage(challenger.ChannelID, "**[Chess]** <@"+challenger.Author.ID+"> You already have a game in progress!")
		return
	}

	opponentID := opponent
	opponentID = strings.TrimLeft(opponentID, "<@!")
	opponentID = strings.TrimRight(opponentID, ">")

	for _, challenge := range challenges {
		if challenger.GuildID == challenge.guildID && (challenger.Author.ID == challenge.challenger || challenger.Author.ID == challenge.opponent) {
			discord.SendChannelMessage(challenger.ChannelID, "**[Chess]** <@"+challenger.Author.ID+"> There's already a challenge by/for you.")
			return
		}

		if challenger.GuildID == challenge.guildID && opponentID == challenge.opponent {
			discord.SendChannelMessage(challenger.ChannelID, "**[Chess]** <@"+challenger.Author.ID+"> That player has already been challenged by someone else.")
			return
		}
	}

	challenge := &challenge{
		guildID:    challenger.GuildID,
		challenger: challenger.Author.ID,
		opponent:   opponentID,
	}

	challenges = append(challenges, challenge)

	if opponentID == s.State.User.ID {
		accept(s.State.User.ID, challenger.GuildID, challenger.ChannelID, s.State.User.ID)
		return
	}

	discord.SendChannelMessage(challenger.ChannelID, "**[Chess]** "+opponent+
		" you have been challenged to a game by <@"+challenger.Author.ID+"> do you accept?")
}

func accept(userID string, guildID string, channelID string, botID string) {
	for index, challenge := range challenges {
		if userID == challenge.opponent && guildID == challenge.guildID {
			createNewGame(index, channelID, botID)
			challenges = append(challenges[:index], challenges[index+1:]...)
			return
		}
	}
}

func decline(m *discordgo.MessageCreate) {
	for index, challenge := range challenges {
		if m.GuildID == challenge.guildID && (m.Author.ID == challenge.challenger || m.Author.ID == challenge.opponent) {
			challenges = append(challenges[:index], challenges[index+1:]...)
			discord.SendChannelMessage(m.ChannelID, "**[Chess]** <@"+m.Author.ID+"> declined the challenge.")
			return
		}
	}
}

func movePiece(m *discordgo.MessageCreate, move string, botID string) {
	var count int
	database.DB.Raw(
		"SELECT IFNULL((SELECT COUNT(id) FROM chess_games WHERE player_white = ? OR player_black = ?), 0) "+
			"FROM chess_games WHERE deleted_at IS NULL",
		m.Author.ID, m.Author.ID).Scan(&count)

	if count == 0 {
		return
	}

	var session chessSession
	database.DB.Raw(
		"SELECT * "+
			"FROM chess_games "+
			"WHERE guild_id = ? AND (player_white = ? OR player_black = ?) AND deleted_at IS NULL "+
			"LIMIT 1",
		m.GuildID, m.Author.ID, m.Author.ID).Scan(&session.model)

	if !(session.model.GameState == models.TurnBlack && session.model.PlayerBlack == m.Author.ID) &&
		!(session.model.GameState == models.TurnWhite && session.model.PlayerWhite == m.Author.ID) {
		return
	}

	pgnReader := strings.NewReader(session.model.BoardState)
	pgn, err := chess.PGN(pgnReader)
	if err != nil {
		slog.Warn("failed to parse PGN data", "error", err)
		discord.SendChannelMessage(m.ChannelID, "**[Chess]** Oops, something went wrong when reading your move. Please try again.")
		return
	}

	session.game = chess.NewGame(pgn)
	err = session.game.MoveStr(move)
	if err != nil {
		slog.Warn("failed to decode chess move or move is invalid", "error", err)
		discord.SendChannelMessage(m.ChannelID, "**[Chess]** <@"+m.Author.ID+">Invalid move! Remember to use algebraic notation!")
		return
	}

	session.updateBoardState()

	if session.isGameOver() {
		session.endGame(m.ChannelID)
		return
	}

	session.updateTurnPlayer()
	session.model.UpdateStates()
	session.sendChannelChessBoard(m.ChannelID)

	if session.isAiTurn(botID) {
		discord.SendChannelMessage(m.ChannelID, "**[Chess]** <@"+botID+"> is thinking about the next move!")
		session.aiMovePiece()
		session.updateBoardState()

		if session.isGameOver() {
			session.endGame(m.ChannelID)
			return
		}

		session.updateTurnPlayer()
		session.model.UpdateStates()
		discord.SendChannelMessage(m.ChannelID, "**[Chess]** <@"+m.Author.ID+"> Your turn to move a piece!")
		session.sendChannelChessBoard(m.ChannelID)
	}
}

func resign(m *discordgo.MessageCreate) {
	var count int
	database.DB.Raw(
		"SELECT IFNULL((SELECT COUNT(id) FROM chess_games WHERE player_white = ? OR player_black = ?), 0) "+
			"FROM chess_games WHERE deleted_at IS NULL",
		m.Author.ID, m.Author.ID).Scan(&count)

	if count == 0 {
		return
	}

	var game *models.ChessGame
	database.DB.Raw(
		"SELECT * "+
			"FROM chess_games "+
			"WHERE guild_id = ? AND (player_white = ? OR player_black = ?) AND deleted_at IS NULL "+
			"LIMIT 1",
		m.GuildID, m.Author.ID, m.Author.ID).Scan(&game)

	if m.Author.ID == game.PlayerWhite {
		game.UpdateGameState(models.WonBlack)
	} else {
		game.UpdateGameState(models.WonBlack)
	}

	database.DB.Delete(game)
	discord.SendChannelMessage(m.ChannelID, "**[Chess]** <@"+m.Author.ID+"> has resigned from the game :(")
}
