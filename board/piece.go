package board

type PieceType byte
type Color byte

const (
	Black Color = 0b0
	White = 0b1
)

const (
        NoPiece PieceType = 0b0000
	Pawn = 0b0010
	Knight = 0b0100
	Bishop = 0b0110
	Rook = 0b1000
	Queen = 0b1010
	King = 0b1100
)

type Color int


type Piece struct {
	data byte
}

func NewPiece(c Color, t PieceType) Piece {
	if c & ~colorMask() != 0 {
		panic("invalid color")
	}

	if t & ~typeMask() != 0 {
		panic("invalid type")
	}

	var cbyte byte = byte(c)
	var tbyte byte = byte(t)
	var data byte = cbyte | tbyte
	return Piece{ data: data }
}

func (p Piece) Color() Color {
	return Color(p.data & colorMask())
}

func (p Piece) Type() PieceType {
	return PieceType(p.data & typeMask())
}

func colorMask() byte {
	return 0b0001
}

func typeMask() byte {
	return ob1110
}
