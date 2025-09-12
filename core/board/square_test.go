package board

import "testing"

func TestBeginSquare(t *testing.T) {
	bs := BeginSqaure()
	if bs.File != BeginFile || bs.Rank != BeginRank {
		t.Errorf("failed to generate begin square")
	}
}

func TestEndSquare(t *testing.T) {
	bs := EndSquare()
	if bs.File != EndFile || bs.Rank != EndRank {
		t.Errorf("failed to generated end square")
	}
}

func TestPanicIfEndShouldNotPanic(t *testing.T) {
	notEnd := Square{
		File:b,
		Rank: 3,
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic at end but it shouldnt")
		}
	}()

	notEnd.panicIfEnd()
}

func TestPanicIfEndShouldPanic(t *testing.T) {
	end := EndSquare()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic at end but it should have")
		}
	}()

	end.panicIfEnd()
}

func TestNextNonEnd(t *testing.T) {
	s := Square{
		File: h,
		Rank: 3,
	}

        s.Next()
        if s.File != a || s.Rank != 4 {
                t.Errorf("failed in next")
        }

        s.Next()
        if s.File != b || s.Rank != 4 {
                t.Errorf("failed in next")
        }
}

func TestNextEnd(t *testing.T) {
        s := Square{
                File: MaxFile,
                Rank: MaxRank,
        }

        s.Next()
        if s != EndSquare() {
                t.Errorf("failed in next end")
        }
}