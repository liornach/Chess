package board

import "piece"

const SQUARES uint = 64
type Square = int

type Board struct {
	board [SQUARES]Piece
}

func NewBoard() Board {
	return Board{}
}


func (b Board) IsOccupied(s Square) bool {
	return b.board[s].Type() != NoPiece
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
	if !IsOccupied(from) {
		return piece.EmptyPiece(), NoPieceAtSquareError{}
	}

	b.board[from] = piece.EmptyPiece()
	return p, nil
}

func (b *Board) Move(from, to Square) error {
	var err error = nil
	var p Piece = piece.EmptyPiece()

	defer func() {
		if err.(SquareOccupiedError) {
			if p == piece.EmptyPiece() {
				panic("fatal error in move (empty piece)")
			}

			if b.Set(p, from) {
				panic("fatal error in move (reset in previous square)")
			}
		}
	}()

	if p, err = b.Take(from); err != nil {
		return err
	}
		
	if err = b.Set(p, to); err != nil {
		return err;
	}

	return nil
}

func (b Board) At(a Square) Piece {
	return b.board[a]
}

type Iterator struct {
	Square Square
	Piece Piece
}

func (b Board) Iterate() <-chan *Iterator {
	c := make(chan *Iterator)

	go func () {
		defer close(c)
		for i := 0; i < SQUARES; i++ [
			c <- &Iterator{Square: i, Piece: b.At(i)}
		]
	}()

	return c
}
