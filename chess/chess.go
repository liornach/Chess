package chess

import "errors"

const squares uint = 64

type piece struct {
	color string
	kind string
}

type board [squares]piece
type file = byte
type rank = uint

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

func (b board) Set(r rank, f file,p piece) {
	assertFile(f)
	assertRank(r)
	idx := fileToIndex(f) + fileToIndex(r)
	b[idx] = // here
}

func (b board) Get(r rank, f file) (piece, bool, error) {
	fidx, err := fileToIndex(f)
	if err != nil {
		return 
	}

	ridx, err := rankToIndex(r)
	if err != nil {

	}
	idx := fileToIndex(f) + rankToIndex(r)
	p := b[idx]
	exist := p != piece{}
	return p, exist
}

func rankFileToIndex(r rank, f file) (uint, error) {
	ridx, err := rankToIndex(r)
	if err != nil {
		return -1, err
	}

	fidx, err := fileToIndex(f)
	if err != nil {
		return -1, err
	}

	ret := fidx + ridx
	return ret, nil
}

func fileToIndex(f file) (uint, error) {
	if err := assertFile(f); err != nil {
		return -1, err
	}

	return uint(f - firstFile())
}

func rankToIndex(r rank) (uint, error) {
	if err := assertRank(r); err != nil {
		return -1, err
	}
	return r - 1 * 8
}

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