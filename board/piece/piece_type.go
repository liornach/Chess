package piece

type PieceType byte


const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

func (pt PieceType) validate() error {
	switch pt {
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