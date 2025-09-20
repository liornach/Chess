package piece

type Color byte

const (
	Black Color = iota
	White
)

func EmptyColor() Color {
	return Color(0b11111111)
}

func (c Color) validate() error {
	switch c {
	case White:
	case Black:
	default:
		return UnknownColorError{}
	}

	return nil
}