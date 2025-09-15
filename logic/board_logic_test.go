package core

import "testing"

func TestFindKing(t *testing.T) {
	b := newBoardState()
	b.Init()
	wk,  err := b.FindKing(White)
	if err != nil || wk.
}
