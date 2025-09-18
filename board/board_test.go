package board

import (
	"testing"
	"example.com/chess/board/piece"
)

var testPiece Piece = piece.WhitePawn()

func TestIsEmptyTrue(t *testing.T) {
	b := Board{}
	res := b.IsEmpty(22)
	exp := true
	if res != exp {
		t.Errorf("exp : %v, res : %v", exp, res)
	}
}

func TestIsEmptyFalse(t *testing.T) {
	b := Board{}
	b.board[23] = &testPiece
	res := b.IsEmpty(23)
	exp := false
	if res != exp {
		t.Errorf("exp : %v, res : %v", exp, res)
	}
}

func TestIsOccupiedTrue(t *testing.T) {
	b := Board{}
	b.board[22] = &testPiece
	res := b.IsOccupied(22)
	exp := true
	if res != exp {
		t.Errorf("exp : %v, res : %v", exp, res)
	}
}

func TestIsOccupiedFalse(t *testing.T) {
	b := Board{}
	res := b.IsOccupied(23)
	exp := false
	if res != exp {
		t.Errorf("exp : %v, res : %v", exp, res)
	}
}

func TestSetNoError(t *testing.T) {
	b := Board{}
	res := b.Set(15 ,testPiece)
	var exp error = nil
	if res != exp {
		t.Errorf("exp : %v, res : %v", exp, res)
	}
}

func TestSetWithError(t *testing.T) {
	b := Board{}
	b.board[16] = &testPiece
	res := b.Set(16, testPiece)
	exp := OccupiedSquareError{}
	if res != exp {
		t.Errorf("exp : %v, res : %v", exp, res)
	}
}

func TestRemoveTrue(t *testing.T) {
	b := Board{}
	b.board[63] = &testPiece
	_, res := b.Remove(63)
	exp := true
	if res != exp {
		t.Errorf("exp : %v, res : %v", exp, res)
	}
}

func TestRemoveFalse(t *testing.T) {
	b := Board{}
	_, res := b.Remove(0)
	exp := false
	if res != exp {
		t.Errorf("exp : %v, res : %v", exp, res)
	}
}