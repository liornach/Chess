package board

type PieceType int

const (
	NoPiece PieceType = iota
	King
	Queen
	Rook
	Knight
	Bishop
	Pawn
)

type Color int

const (
	NoColor Color = iota
	Black
	White
)

type Piece struct {
	Type PieceType
	Color     Color
}

type Board struct {
	board map[Square]Piece
}

func NewBoard() Board {
	return Board{
		board: make(map[Square]Piece),
	}
}

type SquareOccupiedError struct {
}

func (e SquareOccupiedError) Error() string {
	return "sqaure is already occupied"
}

type NoPieceAtSquareError struct {
}

func (e NoPieceAtSquareError) Error() string {
	return "no piece in this square"
}

func (b *Board) Move(from, to Square) error {
	if _, ok := b.board[to]; ok {
		return SquareOccupiedError{}
	}

	if p, err := b.Take(from); err != nil {
		return err
	} else {
		b.board[to] = p
	}

	return nil
}

func (b *Board) Take(from Square) (Piece, error) {
	if p, ok := b.board[from]; !ok {
		return Piece{}, NoPieceAtSquareError{}
	} else {
		delete(b.board, from)
		return p, nil
	}
}

func (b *Board) Set(p Piece, s Square) error {
	if _, ok := b.board[s]; ok {
		return SquareOccupiedError{}
	}

	b.board[s] = p
	return nil
}

func (b Board) At(a Square) (Piece, bool) {
	p, ok := b.board[a]
	return p, ok
}
