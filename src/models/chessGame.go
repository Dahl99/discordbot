package models

import (
	"database/sql"
	"database/sql/driver"

	"discordbot/src/database"
)

type GameState string

const (
	WonWhite  GameState = "WON_WHITE"
	WonBlack  GameState = "WON_BLACK"
	TurnWhite GameState = "TURN_WHITE"
	TurnBlack GameState = "TURN_BLACK"
)

func (c *GameState) Scan(value interface{}) error {
	*c = GameState(value.([]byte))
	return nil
}

func (c GameState) Value() (driver.Value, error) {
	return string(c), nil
}

type ChessGame struct {
	ID          uint64    `json:"id"`
	GuildID     string    `json:"guild-id" gorm:"not null; size:18"`
	PlayerWhite string    `json:"player-white" gorm:"not null; size:18"`
	PlayerBlack string    `json:"player-black" gorm:"not null; size:18"`
	BoardState  string    `json:"board-state" gorm:"not null; type:TEXT"`
	GameState   GameState `json:"game-state" gorm:"type:ENUM('WON_WHITE', 'WON_BLACK', 'TURN_WHITE', 'TURN_BLACK');default:'TURN_WHITE';not null"`
	CreatedAt   int64     `json:"created-at" gorm:"not null"`
	DeletedAt   sql.NullInt64
}

func (c *ChessGame) Update() {
	database.DB.Save(&c)
}
