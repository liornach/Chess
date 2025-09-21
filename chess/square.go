package chess

func rankFileToIndex(r rank, f file) (squareIdx, error) {
	ridx, err := rankToIndex(r)
	if err != nil {
		return invalidSquareIndex(), err
	}

	fidx, err := fileToIndex(f)
	if err != nil {
		return invalidSquareIndex(), err
	}

	ret := fidx + ridx
	return ret, nil
}

func fileToIndex(f file) (squareIdx, error) {
	if err := assertFile(f); err != nil {
		return invalidSquareIndex(), err
	}

	return squareIdx(f - firstFile()), nil
}



func indexToFile(i squareIdx) (file, error) {
	if err := assertSquareIndex(i); err != nil {
		return invalidFile(), err
	}

	ret := firstFile() + file(i & 7)
	return ret, nil
}

func indexToRank(i squareIdx) (rank, error) {
	if err := assertSquareIndex(i); err != nil {
		return invalidRank(), err
	}

	const squares_per_rank squareIdx = 8
	ret := firstRank() + rank(i / squares_per_rank)
	return ret, nil
}

func rankToIndex(r rank) (squareIdx, error) {
	if err := assertRank(r); err != nil {
		return invalidSquareIndex(), err
	}
	return squareIdx(r - 1 * 8), nil
}

func indexToRankFile(i squareIdx) (rank, file, error) {
	if err := assertSquareIndex(i); err != nil {
		return invalidRank(), invalidFile(), err
	}

	ridx , err := indexToRank(i)
	if err != nil {
		return invalidRank(), invalidFile(), err
	}

	fidx, err := indexToFile(i)
	if err != nil {
		return invalidRank(), invalidFile(), err
	}

	return ridx, fidx, nil
}