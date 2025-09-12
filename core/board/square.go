package board

type File int

const (
	BeginFile = a
	a File = iota
	b
	c
	d
	e
	f
	g
	h
	MaxFile = h
	EndFile
)

type Rank int
const BeginRank Rank = 1
const MaxRank = 8
const EndRank = 9

type Square struct {
	File File
	Rank Rank
}

func (s *Square) Next() {
        s.panicIfInvalid()
	s.panicIfEnd()
	
	if s.File == MaxFile {
		if s.Rank == MaxRank {
			*s = EndSquare()
		} else {
			s.File = BeginFile
			s.Rank++
		}

		return
	}

	s.File++
}

func BeginSqaure() Square {
	return Square{
		File: BeginFile,
		Rank: BeginRank,
	}
}

func EndSquare() Square {
	return Square{
		File: EndFile,
		Rank : EndRank,
	}
}

func (s Square) panicIfInvalid() {
        if s.File < BeginFile || s.File > EndFile || s.Rank < BeginRank || s.Rank > EndRank {
                panic("invalid square")
        }
}

func (s Square) panicIfEnd() {
	if s == EndSquare() {
		panic("end square assertion failed")
	}
}
