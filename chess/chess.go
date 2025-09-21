package chess
//C:\Users\Lias1225\sources\go\bin\go.exe

const squares uint = 64



type board [squares]piece
type file = byte
type rank = uint
type squareIdx = int
type color = string
type kind = string

type piece struct {
	color color
	kind kind
}

func invalidRank() rank {
	return 0
}

func invalidFile() file {
	return '\x00'
}

func invalidColor() color {
	return ""
}

func invalidKind() kind {
	return ""
}

func invalidPiece() piece {
	return piece{color: invalidColor(), kind: invalidKind()}
}

func invalidSquareIndex() squareIdx {
	return -1
}

func firstFile() file {
	return 'a'
}

func lastFile() file {
	return 'h'
}

func firstRank() rank {
	return 1
}

func lastRank() rank {
	return 8
}

func firstSquareIdx() squareIdx {
	return 0
}

func lastSquareIdx() squareIdx {
	return int(squares) - 1
}

func (b *board) Set(r rank, f file,p piece) error {
	idx, err := rankFileToIndex(r, f)
	if err != nil {
		return err
	}
	b[idx] = p
	return nil
}

func (b *board) GetByRankFile(r rank, f file) (piece, bool, error) {
	idx, err := rankFileToIndex(r, f)
	if err != nil {
		return invalidPiece(), false, err
	}
	p := b[idx]
	exist := p != piece{}
	return p, exist, nil
}

func (b *board) AtIndex(i squareIdx) (piece, bool, error) {
	if err := assertSquareIndex(i); err != nil {
		return invalidPiece(), false, err
	}

	p := b[i]
	exist := p != piece{}
	return b[i], exist, nil
}

