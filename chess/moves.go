package chess

func (b *board) FindKing(c color) (squareIdx, bool, error) {
	if err := assertColor(c); err != nil {
		return invalidSquareIndex(), false, err
	}

	for i := firstSquareIdx(); i <= lastSquareIdx(); i++ {
		piece, exist, err := b.AtIndex(i)
		if err != nil {
			panic(err) // a bug
		}
		
		if !exist {
			continue
		}

		if piece.kind == "king" && piece.color == c {
			return i, true, nil
		}
	}

	return invalidSquareIndex(), false, nil
}

func (b *board) Moves(i squareIdx) ([]squareIdx, error) {
	piece, exist, err := b.AtIndex(i)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, NoPieceAtSquareError{Square: i}
	}


}

func (b *board) isInDirectSightWithKing(i squareIdx) bool {

}