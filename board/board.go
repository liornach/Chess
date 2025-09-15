package board

const SQUARES uint = 64

type Board struct {
	board [SQUARES]Piece
}

func NewBoard() Board {
	return Board{}
}


func (b Board) IsOccupied(s Square) bool {
	return b.board[s.toInt()].Type() != NoPiece
}


type SquareOccupiedError struct {
}

func (e SquareOccupiedError) Error() string {
	return "sqaure is already occupied"
}


func (b *Board) Set(p Piece, s Square) error {
	if b.IsOccupied(s) {
		return SquareOccupiedError{}
	}

	b.board[s] = p
	return nil
}

type NoPieceAtSquareError struct {
}

func (e NoPieceAtSquareError) Error() string {
	return "no piece in this square"
}

func (b *Board) Take(from Square) (Piece, error) {
	if p, ok := b.board[from]; !ok {
		return Piece{}, NoPieceAtSquareError{}
	} else {
		delete(b.board, from)
		return p, nil
	}
}

func (b *Board) Move(from, to Square) error {
	if b.IsOccupied(to) {
		return SquareOccupiedError{}
	}

	if p, err := b.Take(from); err != nil {
		return err
	} else {
		b.board[to] = p
	}

	return nil
}

func (b Board) At(a Square) (Piece, bool) {
	p, ok := b.board[a]
	return p, ok
}

type Iterator struct {
	Square Square
	Piece Piece
}

func (b Board) Iterate() <-chan *Iterator {
	c := make(chan *Iterator)

	go func () {
		defer close(c)
		for square, piece := range b.board {
			c <- &Iterator{Square: square, Piece: piece}
		}
	}()

	return c
}
