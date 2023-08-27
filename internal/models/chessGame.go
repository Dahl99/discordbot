package models

import (
	"database/sql/driver"
	"github.com/Dahl99/DiscordBot/internal/database"
	"time"

	"gorm.io/gorm"
)

type GameState string

const (
	WonWhite  GameState = "WON_WHITE"
	WonBlack  GameState = "WON_BLACK"
	TurnWhite GameState = "TURN_WHITE"
	TurnBlack GameState = "TURN_BLACK"
	Draw      GameState = "DRAW"
)

func (c GameState) Value() (driver.Value, error) {
	return string(c), nil
}

func (c *GameState) Scan(value interface{}) error {
	*c = GameState(value.([]byte))
	return nil
}

type ChessGame struct {
	ID          uint64         `json:"id" gorm:"primaryKey"`
	GuildID     string         `json:"guildId" gorm:"not null; size:18"`
	PlayerWhite string         `json:"playerWhite" gorm:"not null; size:18"`
	PlayerBlack string         `json:"playerBlack" gorm:"not null; size:18"`
	BoardState  string         `json:"boardState" gorm:"not null; type:TEXT"`
	GameState   GameState      `json:"gameState" gorm:"type:ENUM('WON_WHITE', 'WON_BLACK', 'TURN_WHITE', 'TURN_BLACK', 'DRAW');default:'TURN_WHITE';not null"`
	CreatedAt   *time.Time     `json:"createdAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

func (c *ChessGame) UpdateStates() {
	database.DB.Model(&c).Select("board_state", "game_state").Updates(ChessGame{BoardState: c.BoardState, GameState: c.GameState})
}

func (c *ChessGame) UpdateGameState(gameState GameState) {
	database.DB.Model(&c).Update("game_state", gameState)
}
