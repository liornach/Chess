package board

type BoardBuilder struct {
	board [SQUARES]Piece
	i int
}

func (b *BoardBuilder) Build() Board {
	return Board{
		board: b.board
	}
}

func (b *BoardBuilder) Put(p Piece) *BoardBuilder {
	if i == SQUARES {
		panic("out of bounds access to board")
	}

	b.board[i] = p
	i++
	return b
}

func (b *BoardBuilder) PutRow(p Piece) *BoardBuilder {
	const ROW_SQUARES := 8
	for i := 0; i < ROW_SQUARES; i++ {
		b.Put(p)
	}

	return b
}
