package piece

func WhitePawn() Piece {
	return whitePiece(Pawn)
}

func WhiteKnight() Piece {
	return whitePiece(Knight)
}

func WhiteBishop() Piece {
	return whitePiece(Bishop)
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

func whitePiece(t PieceType) Piece {
	return Piece{ptype : t, color : White}
}

func BlackPawn() Piece {
	return blackPiece(Pawn)
}

func BlackKnight() Piece {
	return blackPiece(Knight)
}

func BlackBishop() Piece {
	return blackPiece(Bishop)
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

func blackPiece(t PieceType) Piece {
	return Piece{ptype : t, color : White}
}

