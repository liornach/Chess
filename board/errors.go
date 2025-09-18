package board

type OccupiedSquareError struct {
}

func (e OccupiedSquareError) Error() string {
	return "square is occupied"
}

