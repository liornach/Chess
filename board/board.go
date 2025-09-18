package board

import (
	"example.com/chess/board/piece"
)

type Piece = piece.Piece
type Square = uint
const SQUARES uint = 64

type Board struct {
	board [SQUARES]*Piece
}

func (b *Board) Set(s Square, p Piece) error {
	if b.IsOccupied(s) {
		return OccupiedSquareError{}
	}

	b.board[s] = &p
	return nil
}

func (b *Board) Remove(s Square) (*Piece, bool) {
	if b.IsEmpty(s) {
		return piece.EmptyPiece(), false
	}
	
	p := b.board[s]
	b.board[s] = piece.EmptyPiece()
	return p, true
}

func (b *Board) IsEmpty(s Square) bool {
	res := b.board[s] == piece.EmptyPiece()
	return res
}

func (b *Board) IsOccupied(s Square) bool {
	return !b.IsEmpty(s)
}