package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"

	"discordbot/src/database"
)

type GameState string

const (
	WonWhite  GameState = "WON_WHITE"
	WonBlack  GameState = "WON_BLACK"
	TurnWhite GameState = "TURN_WHITE"
	TurnBlack GameState = "TURN_BLACK"
	Draw      GameState = "DRAW"
)

func (c *GameState) Scan(value interface{}) error {
	*c = GameState(value.([]byte))
	return nil
}

func (c GameState) Value() (driver.Value, error) {
	return string(c), nil
}

type ChessGame struct {
	ID          uint64    `gorm:"primaryKey"`
	GuildID     string    `gorm:"not null; size:18"`
	PlayerWhite string    `gorm:"not null; size:18"`
	PlayerBlack string    `gorm:"not null; size:18"`
	BoardState  string    `gorm:"not null; type:TEXT"`
	GameState   GameState `gorm:"type:ENUM('WON_WHITE', 'WON_BLACK', 'TURN_WHITE', 'TURN_BLACK', 'DRAW');default:'TURN_WHITE';not null"`
	CreatedAt   time.Time `gorm:"not null"`
	DeletedAt   gorm.DeletedAt
}

func (c *ChessGame) Update() {
	database.DB.Save(&c)
}
