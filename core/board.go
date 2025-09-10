package core

// C:\Users\Lias1225\sources\go\bin 

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
	pieceType PieceType
	color Color
}

type File int
const (
	a File = iota
	b
	c
	d
	e
	f
	g
	h
	MaxFile
)

const MinRank = 1
const MaxRank = 8

type Square int

func beginSquare() Square {
	return newSquare(1, a)
}

func endSquare() Square {
	return newSquare(8, MaxFile)
}

func newSquare(rank int, file File) Square {
	return (Square)((rank - 1) * 8 + (int)(file))
}

func (s Square) File() File {
	return (File)(s % 8)
}

func (s Square) Rank() int {
	return (int)(s / 8)
}

type Board struct {
	board map[Square]Piece
}

type squareOccupiedError struct {
}

func (e squareOccupiedError) Error() string {
	return "sqaure is already occupied"
}

type noPieceAtSquareError struct {
}

func (e noPieceAtSquareError) Error() string {
	return "no piece in this square"
}

func (b *Board) Move(from, to Square) error {
	if _, ok := b.board[to]; ok {
		return squareOccupiedError{}
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
		return Piece{}, noPieceAtSquareError{}
	} else {
		delete(b.board, from)
		return p, nil
	}
}

func (b *Board) Set(p Piece, s Square) error {
	if _, ok := b.board[s]; ok {
		return squareOccupiedError{}
	}

	b.board[s] = p
	return nil
}

func (b Board) At(a Square) (Piece, bool) {
	p, ok := b.board[a]
	return p, ok
}