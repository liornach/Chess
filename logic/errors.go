package logic

type NoPieceAtSquareError struct {
}

func (e NoPieceAtSquareError) Error() string {
	return "no piece at square"
}

type TurnViolationError struct {
}

func (e TurnViolationError) Error() string {
	return "this is not this color's turn"
}