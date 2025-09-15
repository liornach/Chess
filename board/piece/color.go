package piece

type Color byte

const (
	NoColor Color = iota
	Black
	White
)

func (c Color) validate() error {
	switch c {
	case NoColor:
	case White:
	case Black:
	default:
		return UnknownColorError{}
	}

	return nil
}