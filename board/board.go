package board

import (
	"example.com/chess/board/piece"
	"example.com/chess/board/square"
)

type Piece = piece.Piece
type Square = square.Square
const SQUARES uint = 64

type Board struct {
	board [SQUARES]Piece
}

func (b *Board) Move(from, to Square) error {
	if b.IsEmpty(from) {
		return NoPieceAtSquareError{}
	}

	if b.IsOccupied(to) {
		return OccupiedSquareError{}
	}
}

func (b *Board) remove(s Square) Piece {
	idx, err := s.ToUint()
	if err != nil {
		panic(err)
	}

	p := b[idx]
	b[idx] = piece.EmptyPiece()
	return p
}

func (b *Board) IsEmpty(s Square) bool {
	si, err := s.ToUint()
	if err != nil {
		panic(err)
	}

	return b.board[si].Type() != piece.NoPiece
}

func (b *Board) IsOccupied(s Square) bool {
	return !IsFree(s)
}