package square

type InvalidRankError struct {
}

func (e InvalidRankError) Error() string {
	return "invalid rank"
}

type InvalidFileError struct {
}

func (e InvalidFileError) Error() string {
	return "invalid file"
}