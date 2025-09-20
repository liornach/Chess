package chess

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

func (b board) Get(r rank, f file) (piece, bool) {
	assertFile(f)
	assertRank(r)
	idx := fileToIndex(f) + rankToIndex(r)
	p := b[idx]
	exist := p != piece{}
	return p, exist
}

func fileToIndex(f file) uint {
	assertFile(f)
	return uint(f - firstFile())
}

func rankToIndex(r rank) uint {
	assertRank(r)
	return r - 1 * 8
}

func assertFile(f file) {
	if f < firstFile() || f > lastFile() {
		panic("file is out of range")
	}
}

func assertRank(r rank) {
	if r < firstRank() || r > lastRank() {
		panic("rank is out of range")
	}
}