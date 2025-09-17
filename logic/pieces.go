package logic

import (
	"example.com/chess/board/piece"
	"example.com/chess/board/piece/factory"

)

type Piece = piece.Piece

var whiteRook Piece
var whiteKnight Piece
var whiteBishop Piece
var whiteQueen Piece
var whiteKing Piece
var whitePawn Piece
var blackRook Piece
var blackKnight Piece
var blackBishop Piece
var blackQueen Piece
var blackKing Piece
var blackPawn Piece

func init() {
	whitePawn = factory.WhitePawn()
	whiteRook = factory.WhiteRook()
	whiteKnight = factory.WhiteKnight()
	whiteBishop = factory.WhiteBishop()
	whiteQueen = factory.WhiteQueen()
	whiteKing = factory.WhiteKing()
	blackPawn = factory.BlackPawn()
	blackRook = factory.BlackRook()
	blackKnight = factory.BlackKnight()
	blackBishop = factory.BlackBishop()
	blackQueen = factory.BlackQueen()
	blackKing = factory.BlackKing()	
}