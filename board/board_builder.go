package board

type BoardBuilder struct {
	board [SQUARES]Piece
	i uint
	ri uint
}

const squares_in_row = 8

func NewBoardBuilder() BoardBuilder {
	return BoardBuilder{
		i = 0
		ri = 7
	}
}

func (b *BoardBuilder) Build() Board {
	return Board{
		board: b.board,
	}
}

func (b *BoardBuilder) Put(p Piece) *BoardBuilder {
	if ri == i {
		i = (i + squares_in_row) & ~squares_in_row
	}

	if b.i == SQUARES {
		panic("out of bounds access to board")
	}

	b.board[b.i] = p
	b.i++

	return b
}

func (b *BoardBuilder) putRev(p Piece) *BoardBuilder {
	
}

func (b *BoardBuilder) PutMirror(p Piece) *BoardBuilder {
	Put(p)
	b.board[ri] = p
	ri--

	return b
}

func (b *BoardBuilder) PutRow(p Piece) *BoardBuilder {
	const ROW_SQUARES = 8
	for i := 0; i < ROW_SQUARES; i++ {
		b.Put(p)
	}

	return b
}