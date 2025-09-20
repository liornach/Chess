package logic

import (
	"example.com/chess/board"
	"example.com/chess/board/piece"
	"example.com/chess/logic/square"
)

type Board = board.Board
type Color = piece.Color
type Square = square.Square

type turnStateBoard struct {
	board Board
	turn Color
}

func (t *turnStateBoard) Move(from, to Square) error {
	if fu , err := from.ToUint(); err != nil {
		return err
	} else if t.board.IsEmpty(fu) {
		return NoPieceAtSquareError{}
	} else if tu, err := to.ToUint(); err != nil {
		return err 
	} else if t.board.IsOccupied(tu) {
		return board.OccupiedSquareError{}
	} else if p, ok := t.board.Remove(fu); !ok {
		panic("failed to remove in move")
	}  else if err := t.board.Set(tu, *p); err != nil {
		panic(err) // should never occur
	}

	return nil
}