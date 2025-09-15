package square

type Rank byte

func (r Rank) validate() error {
	switch byte(r) {
	case 'a':
	case 'b':
	case 'c':
	case 'd':
	case 'e':
	case 'f':
	case 'g':
	case 'h':
	default:
		return InvalidRankError{}
	}

	return nil
}