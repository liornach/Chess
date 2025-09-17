package board

import "piece"

type BoardBuilder struct {
	square int
	board [SQUARES]Piece
}

func (bb *BoardBuilder) SetNext(p Piece) bool {
	bb.board[bb.square] = p
	bb.square++
	return bb.square == SQUARES
}

func (bb *BoardBuilder) SetAt(p Piece, i int) {
	bb.board[i] = p
}

func (bb *BoardBuilder) Build() Board {
	return Board{
		board: bb.board
	}
}
