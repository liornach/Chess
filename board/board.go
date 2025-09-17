package board

import (
	"example.com/chess/board/piece"
)

type Piece = piece.Piece
const SQUARES uint = 64

type Board struct {
	board [SQUARES]Piece
}
