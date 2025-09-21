package chess

import "fmt"

type FileOutOfRangeError struct {
	File file
}

func (e FileOutOfRangeError) Error() string {
	return fmt.Sprintf("file %v is out of range", e.File)
}

type RankOutOfRangeError struct {
	Rank rank
}

func (e RankOutOfRangeError) Error() string {
	return fmt.Sprintf("rank %v is out of range", e.Rank)
}

type NoPieceAtSquareError struct {
	Square squareIdx
}

func (e NoPieceAtSquareError) Error() string {
	return fmt.Sprintf("the is no piece at square idx %v", e.Square)
}

type PieceNotFoundError struct {
	Piece piece
}

func (e PieceNotFoundError) Error() string {
	return fmt.Sprintf("piece %v not found", e.Piece)
}

type InvalidColorError struct {
	Color color
}

func (e InvalidColorError) Error() string {
	return fmt.Sprintf("invalid color - %s", e.Color)
}

type InvalidKindError struct {
	Kind kind
}

func (e InvalidKindError) Error() string {
	return fmt.Sprintf("invalid kind - %s", e.Kind)
}

type InvalidSquareIndexError struct {
	SquareIndex squareIdx
}

func (e InvalidSquareIndexError) Error() string {
	return fmt.Sprintf("invalid square index - %d", e.SquareIndex)
}