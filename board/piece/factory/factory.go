package factory

import "example.com/chess/board/piece"

type Piece = piece.Piece
type PieceType = piece.PieceType
const White = piece.White
const Black = piece.Black
const Pawn = piece.Pawn
const Bishop = piece.Bishop
const Knight = piece.Knight
const Rook = piece.Rook
const Queen = piece.Queen
const King = piece.King

func WhitePawn() Piece {
	return whitePiece(Pawn)
}

func WhiteBishop() Piece {
	return whitePiece(Bishop)
}

func WhiteKnight() Piece {
	return whitePiece(Knight)
}

func WhiteRook() Piece {
	return whitePiece(Rook)
}

func WhiteQueen() Piece {
	return whitePiece(Queen)
}

func WhiteKing() Piece {
	return whitePiece(King)
}

func BlackPawn() Piece {
	return blackPiece(Pawn)
}

func BlackBishop() Piece {
	return blackPiece(Bishop)
}

func BlackKnight() Piece {
	return blackPiece(Knight)
}

func BlackRook() Piece {
	return blackPiece(Rook)
}

func BlackQueen() Piece {
	return blackPiece(Queen)
}

func BlackKing() Piece {
	return blackPiece(King)
}

func EmptyPiece() Piece {
	return piece.EmptyPiece()
}

func whitePiece(t PieceType) Piece {
	p, err := piece.NewPiece(White, t)
	if err != nil {
		panic(err)
	}

	return p
}

func blackPiece(t PieceType) Piece {
	p, err := piece.NewPiece(Black, t)
	if err != nil {
		panic(err)
	}

	return p
}