package board

import (
	"example.com/chess/board/piece"
	"example.com/chess/board/piece/factory"
)

type Piece = piece.Piece
const SQUARES uint = 64

type Board struct {
	board [SQUARES]Piece
}

func NewBoard() Board {{
	}
}
