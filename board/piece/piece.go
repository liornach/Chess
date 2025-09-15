package piece


type Piece struct {
	data byte
}

func NewPiece(c Color, t PieceType) (Piece, error) {
	if err := c.validate(); err != nil {
		return Piece{}, err
	}

	if err := t.validate(); err != nil {
		return Piece{}, err
	}


	var cbyte byte = byte(c)
	var tbyte byte = byte(t)
	var data byte = cbyte | tbyte
	return Piece{ data: data }, nil
}

func (p Piece) Color() Color {
	return Color(p.data & colorMask())
}

func (p Piece) Type() PieceType {
	return PieceType(p.data & typeMask())
}