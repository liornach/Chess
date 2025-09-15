package piece

type PieceType byte


const (
    NoPiece PieceType = 0b0
	Pawn PieceType = iota + 4
	Knight
	Bishop
	Rook
	Queen
	King
)

func (pt PieceType) validate() error {
	switch pt {
	case NoPiece:
	case Pawn:
	case Knight:
	case Bishop:
	case Rook:
	case Queen:
	case King:
		return nil
	}

	return UnknownPieceTypeError{}
}

func typeMask() byte {
	return 0b1100
}