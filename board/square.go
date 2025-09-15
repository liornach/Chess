package board

type File byte

const (
	a File = iota + 9
	b
	c
	d
	e
	f
	g
	h
)


type Rank byte
const (
	FirstRank Rank = 1
	LastRank 8
)

type Square struct {
	data byte
}

func NewSquare(f File, r Rank) {
	if f < a || f > h {
		panic("invalid file")
	}

	if r < FirstRank || r > LastRank {
		panic("invalid rank")
	}

	return Square {
		data: byte(r) | byte(f)
	}
}

//func (s Square) Next() (Square, error) {
//	intSquare, err  := s.toInt()
//        if err != nil {
//                return Square{}, err
//        }
//
//        return fromInt(intSquare + 1)
//}
//

func fromInt(i int) (Square, error) {
        if err := assert(i); err != nil {
                return Square{}, err
        }

        file := (i - 1 ) & (8 - 1) + 1
        rank := (i - 1) / 8 + 1
        return Square{
                File: File(file),
                Rank: Rank(rank),
        }, nil
}

func (s Square) toInt() (int, error) {
        const FILES_PER_RANK int = 8
        res := int(s.File) + int(s.Rank - 1) * FILES_PER_RANK
        return res, assert(res)
}

func lastSquareInt() int {
        const SQUARES int = 64
        return SQUARES
}

func firstSquareInt() int {
        const FIRST_SQUARE int = 1
        return FIRST_SQUARE
}

func FirstSquare() Square {
        first, _ := fromInt(firstSquareInt())
        return first
}

func LastSquare() Square {
        last, _ := fromInt(lastSquareInt())
        return last
}

type InvalidSquareError struct{
}

func (e InvalidSquareError) Error() string {
        return "invalid square"
}

func assert(i int) error {
        if i < firstSquareInt() || i > lastSquareInt() {
                return InvalidSquareError{}
        }

        return nil
}
