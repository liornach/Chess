package core

//import "example.com/chess/board"

type BoardState struct {
	board board
	turn  Color
}

func NewBoardState() BoardState {
	return BoardState{
		board: newBoard(),
		turn:  0,
	}
}

func (b *BoardState) setOrPanic(p Piece, s Square) {
	if err := b.board.Set(p, s); err != nil {
		panic(err)
	}
}

const WHITE_PIECES_RANK int = 1
const BLACK_PIECES_RANK int = 8

func (b *BoardState) setDoublePieces(pt PieceType, fileOffset File) {
	white := Piece{pieceType: pt, color: White}
	black := Piece{pieceType: pt, color: Black}
	leftFile := h - fileOffset
	rightFile := a + fileOffset
	b.setOrPanic(white, newSquare(WHITE_PIECES_RANK, leftFile))
	b.setOrPanic(white, newSquare(WHITE_PIECES_RANK, rightFile))
	b.setOrPanic(black, newSquare(BLACK_PIECES_RANK, leftFile))
	b.setOrPanic(black, newSquare(BLACK_PIECES_RANK, rightFile))
}

func (b *BoardState) Init() {
	whitePawn := Piece{pieceType: Pawn, color: White}
	blackPawn := Piece{pieceType: Pawn, color: Black}
	const WHITE_PAWN_RANK int = 2
	const BLACK_PAWN_RANK int = 7
	for f := a; f < MaxFile; f++ {
		b.setOrPanic(whitePawn, newSquare(WHITE_PAWN_RANK, f))
		b.setOrPanic(blackPawn, newSquare(BLACK_PAWN_RANK, f))
	}

	b.setDoublePieces(Rook, 0)
	b.setDoublePieces(Knight, 1)
	b.setDoublePieces(Bishop, 2)

	b.setOrPanic(Piece{pieceType: King, color: White}, newSquare(WHITE_PIECES_RANK, e))
	b.setOrPanic(Piece{pieceType: King, color: Black}, newSquare(BLACK_PIECES_RANK, e))

	b.setOrPanic(Piece{pieceType: Queen, color: White}, newSquare(WHITE_PIECES_RANK, d))
	b.setOrPanic(Piece{pieceType: Queen, color: Black}, newSquare(BLACK_PIECES_RANK, d))

	b.turn = White
}

type PieceNotFoundError struct {
}

func (e PieceNotFoundError) Error() string {
	return "piece could not be found"
}

func (b *BoardState) FindKing(c Color) (Square, error) {
	king := Piece{ Type: King, Color : c}

	for iter := range b.board.Iterate() {
		if iter.Piece == king {
			return iter.Square, nil
		}
	}

	return Square{}, PieceNotFoundError{}
}
