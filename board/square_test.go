package board

import "testing"

func TestAssert(t *testing.T) {
        valid := 25
        if err := assert(valid); err != nil {
                t.Errorf("failed on valid assertion")
        }

        invalid := -1
        if err := assert(invalid); err == nil {
                t.Errorf("failed on invalid assertion")
        } 

        invalid = 65
        if err := assert(invalid); err == nil {
                t.Errorf("failed on invalid assertion")
        }
}

func TestFirstSquare(t *testing.T) {
	bs := FirstSquare()
	if bs.File != a || bs.Rank != Rank(1) {
		t.Errorf("failed to generate first square")
	}
}

func TestLastSquare(t *testing.T) {
	bs := LastSquare()
	if bs.File != h || bs.Rank != Rank(8) {
		t.Errorf("failed to generated last square")
	}
}


func TestValidNext(t *testing.T) {
	s := Square{
		File: h,
		Rank: 3,
	}

        var err error 
        if s, err = s.Next(); err != nil || s.File != a || s.Rank != 4 {
                t.Errorf("failed in next")
        }

        if s, err = s.Next(); err != nil || s.File != b || s.Rank != 4 {
                t.Errorf("failed in next")
        }
}

func TestInvalidNext(t *testing.T) {
        s := LastSquare()
        if _, err := s.Next(); err == nil {
                t.Errorf("failed in next")
        }
}