package piece


type Piece struct {
	ptype PieceType
	color Color
}

func NewPiece(c Color, t PieceType) (Piece, error) {
	if err := c.validate(); err != nil {
		return Piece{}, err
	}

	if err := t.validate(); err != nil {
		return Piece{}, err
	}

	return Piece{ ptype : t, color : c }, nil
}

func (p Piece) Color() Color {
	return Color(p.color)
}

func (p Piece) Type() PieceType {
	return PieceType(p.ptype)
}