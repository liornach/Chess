package chess

func assertRankFile(r rank, f file) error {
	if err := assertRank(r); err != nil {
		return err
	}

	if err := assertFile(f); err != nil {
		return err
	}

	return nil
}

func assertFile(f file) error {
	if f < firstFile() || f > lastFile() {
		return FileOutOfRangeError{File: f}
	}

	return nil
}

func assertRank(r rank) error {
	if r < firstRank() || r > lastRank() {
		return RankOutOfRangeError{Rank :r}
	}

	return nil
}

func assertColor(c color) error {
	if c != "white" && c != "black" {
		return InvalidColorError{Color : c}
	}

	return nil
}

func assertKind(k kind) error {
	switch k {
	case "pawn":
	case "bishop":
	case "knight":
	case "rook":
	case "queen":
	case "king":
	default:
		return InvalidKindError{Kind: k}
	}

	return nil
}

func assertSquareIndex(i squareIdx) error {
	if i < firstSquareIdx() || i > lastSquareIdx() {
		return InvalidSquareIndexError{SquareIndex: i}
	}

	return nil
}