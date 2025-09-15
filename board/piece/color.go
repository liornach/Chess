package piece

type Color byte

const (
	NoClor Color = iota
	Black
	White
)

func (c Color) validate() error {
	switch c {
	case White:
	case Black:
	default:
		return UnknownColorError{}
	}

	return nil
}

func colorMask() byte {
	return 0b0011
}