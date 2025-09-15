package piece

type UnknownPieceTypeError struct {
}

func (e UnknownPieceTypeError) Error() string {
	return "unknown piece type"
}

type UnknownColorError struct {
}

func (e UnknownColorError) Error() string {
	return "unknown color"
}