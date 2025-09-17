package logic

import "example.com/chess/board"
//import "example.com/chess/board/piece"

type Board = board.Board
type BoardBuilder = board.BoardBuilder

type Logic struct {
	board Board
}

func NewLogic() Logic {
	bb := BoardBuilder{}
	bb.PutTwo(whiteRook).PutTwo(whiteKnight).PutTwo(whiteBishop).Put(whiteQueen).Put(whiteKing)

	return Logic{
		board: bb.Build(),
	}
}