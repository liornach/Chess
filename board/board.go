package board

import (
	"example.com/chess/board/piece"
	"example.com/chess/board/square"
)

type Piece = piece.Piece
const SQUARES uint = 64

type Board struct {
	board [SQUARES]Piece
}

func (b *Board) Move(from, to Square)