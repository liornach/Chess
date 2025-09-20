package square

type Square struct {
	rank Rank
        file File
}

func NewSquare(r byte, f int) (Square, error) {
        file := File(f)
        if err := file.validate(); err != nil {
                return Square{}, err
        }

        rank := Rank(r)
        if err := rank.validate(); err != nil {
                return Square{}, err
        }

	return Square{
		rank: rank,
                file: file,
	}, nil
}

func fromInt(i int) (Square, error) {
        if err := assert(i); err != nil {
                return Square{}, err
        }

        file := (i) & (8 - 1)
        rank := (i) / 8
        return Square{
                rank: Rank(rank) + firstRank(),
                file: File(file) + firstFile(),
        }, nil
}

func firstFile() File {
        return File('a')
}

func firstRank() Rank {
        return Rank(1)
}

func (s Square) ToUint() (uint, error) {
        const FILES_PER_RANK int = 8
        res := int(s.file) + int(s.rank - 1) * FILES_PER_RANK
        return uint(res), assert(res)
}

func lastSquareInt() int {
        const SQUARES int = 63
        return SQUARES
}

func firstSquareInt() int {
        const FIRST_SQUARE int = 0
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
