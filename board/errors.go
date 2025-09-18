package board

type NoPieceAtSquareError struct {
}

func (e *NoPieceAtSquareError) Error() {
	return "no piece in square"
}

type OccupiedSquareError struct {
}

func (e *OccupiedSquareError) Error() {
	return "square is occupied"
}

