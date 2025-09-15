package piece

import "testing"

func TestValidatePawn(t *testing.T) {
	p := Pawn
	res := p.validate()
	if res != nil {
		t.Errorf("failed with err : %v", res.Error())
	}
}

func TestValidateBishop(t *testing.T) {
	p := Bishop
	res := p.validate()
	if res != nil {
		t.Errorf("failed with err : %v", res.Error())
	}
}

func TestValidateKing(t *testing.T) {
	p := King
	res := p.validate()
	if res != nil {
		t.Errorf("failed with err : %v", res.Error())
	}
}

func TestValidateKnight(t *testing.T) {
	p := Knight
	res := p.validate()
	if res != nil {
		t.Errorf("failed with err : %v", res.Error())
	}
}

func TestValidateQueen(t *testing.T) {
	p := Queen
	res := p.validate()
	if res != nil {
		t.Errorf("failed with err : %v", res.Error())
	}
}

func TestValidateRook(t *testing.T) {
	p := Rook
	res := p.validate()
	if res != nil {
		t.Errorf("failed with err : %v", res.Error())
	}
}

func TestValidateNoPiece(t *testing.T) {
	p := NoPiece
	res := p.validate()
	if res != nil {
		t.Errorf("failed with err : %v", res.Error())
	}
}

func TestValidatePieceExpectSomeError(t *testing.T) {
	p := PieceType(0b11111111)
	res := p.validate()
	if res == nil {
		t.Errorf("failed, expected some error")
	}
}