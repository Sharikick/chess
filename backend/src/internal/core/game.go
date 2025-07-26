// Package core ...
package core

import (
	"github.com/google/uuid"
)

type Game struct {
	ID  uuid.UUID
	FEN string
}

func InitFEN() string {
	return ""
}

func NewGame() *Game {
	return &Game{
		ID:  uuid.New(),
		FEN: InitFEN(),
	}
}

type Board [8][8]byte
