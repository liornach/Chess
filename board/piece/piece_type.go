package piece

type PieceType byte


const (
    NoPiece PieceType = iota
	Pawn
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
	default:
		return UnknownPieceTypeError{}
	}

	return nil
}