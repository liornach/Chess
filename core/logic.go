package core

type GameState struct {
	Board Board
	Turn Color
}

func (b *Board) setOrPanic(p Piece, s Square) {
	if err := g.Board.Set(p, s); err != nil {
		panic(err)
	}
}

func (g *GameState) setDoublePieces(pt PieceType, fileOffset File) {
	const WHITE_PIECES_RANK int = 1
	const BLACK_PIECES_RANK int = 8
	white := Piece{pieceType: pt, color : White}
	black := Piece{pieceType: pt, color : Black}
	leftFile := h - fileOffset
	rightFile := a + fileOffset
	setOrPanic(white, newSquare(WHITE_PIECES_RANK, leftFile))
	setOrPanic(white, newSquare(WHITE_PIECES_RANK, rightFile))
	setOrPanic(black, newSquare(BLACK_PIECES_RANK, leftFile))
	setOrPanic(black, newSquare(BLACK_PIECES_RANK, rightFile))
}


func (g *GameState) Init() {
	whitePawn := Piece{pieceType: Pawn, color : White}
	blackPawn := Piece{pieceType: Pawn, color : Black}
	const WHITE_PAWN_RANK int = 2
	const BLACK_PAWN_RANK int = 7
	for f := a; f < MaxFile; f++ {
		setOrPanic(whitePawn, newSquare(WHITE_PAWN_RANK, f))
		setOrPanic(blackPawn, newSquare(BLACK_PAWN_RANK, f))
	}

	setDoublePieces(Rook, 0)
	setDoublePieces(Knight, 1)
	setDoublePieces(Bishop, 2)
}

type Logic struct {
	Board Board
}