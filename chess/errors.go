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
