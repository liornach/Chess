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

func NewBoard() Board {
	var board [SQUARES]Piece
	board[0] = factory.WhiteRook()
	board[1] = factory.WhiteKnight()
	board[2] = factory.WhiteBishop()
	board[3] = factory.WhiteQueen()
	board[4] = factory.WhiteKing()
	board[5] = factory.WhiteBishop()
	board[6] = factory.WhiteKnight()
	board[7] = factory.WhiteRook()

	const PAWNS int = 8
	const FIRST_PAWN_SQUARE = 8
	for i := 0; i < PAWNS; i++ {
		board[FIRST_PAWN_SQUARE + i] = factory.WhitePawn()
	}

	const EMPTY_SQUARES int = 32
	const FIRST_EMPTY_SQUARE = 16
	for i := 0; i < EMPTY_SQUARES; ++i {
		board[FIRST_EMPTY_SQUARE + i] = factory.EmptyPiece()
	}
}
