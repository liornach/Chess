package chess

func up(i SquareIndex) (squareIdx, error) {
	rank, file, err := indexToRankFile(i)
	if err != nil {
		return i,err
	}

	rank++
	idx, err := rankFileToIndex(rank, file)
	if err != nil {
		return i, err
	}

	return idx, nil
}

func down(i squareIdx) (squareIdx, error) {
	rank, file, err := indexToRankFile(i)
	if err != nil {
		return i, err
	}

	rank--
	idx, err := rankFileToIndex(rank, file)
	if err != nil {
		return i, err
	}

	return idx, nil
}

func right(i squareIdx) (squareIdx, error) {
	rank, file, err := indexToRankFile(i)
	if err != nil {
		return i, err
	}

	file++
	idx, err := rankFileToIndex(rank, file)	
	if err != nil {
		return i, err
	}

	return idx, nil
}

func left(i squareIdx) (squareIdx, error) {
	rank, file, err := indexToRankFile(i)
	if err != nil {
		return i, err
	}

	file--
	idx, err := rankFileToIndex(rank, file)	
	if err != nil {
		return i, err
	}

	return idx, nil
}

func upRight(i squareIdx) (squareIdx, error) {
	right, err := right(i)
	if err != nil {
		return i, err
	}

	upright, err := up(right)
	if err != nil {
		return i, err
	}

	return upright, nil
}

func downRight(i squareIdx) (squareIdx, error) {
	right, err := right(i)
	if err != nil {
		return i, err
	}

	downRight, err := down(right)
	if err != nil {
		return i, err
	}

	return downRight, nil
}

func upLeft(i squareIdx) (squareIdx, error) {
	l, err := left(i)
	if err != nil {
		return i, err
	}

	ul, err := up(l)
	if err != nil {
		return i, err
	}

	return ul, nil
}

func downLeft(i squareIdx) (squareIdx, error) {
	l, err := left(i)
	if err != nil {
		return i, err
	}

	dl, err := down(l)
	if err != nil {
		return i, err
	}

	return dl, nil
}

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