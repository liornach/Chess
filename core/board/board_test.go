package board

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	if b.board == nil {
		t.Errorf("failed in new board")
	}
}

func TestIsOccupied(t *testing.T) {
	b := NewBoard()
	s := FirstSquare()
	if b.IsOccupied(s) {
		t.Errorf("failed in is occupied")
	}

	b.board[s] = Piece{}
	if !b.IsOccupied(s) {
		t.Errorf("failed in is occupied")
	}

	delete(b.board, s)
	if b.IsOccupied(s) {
		t.Errorf("failed in is occupied")
	}
}

func TestSet(t *testing.T) {
        b := NewBoard()
        s := FirstSquare()
        s, _ = s.Next()

        if b.Set(Piece{}, s) != nil {
                t.Errorf("failed in set")
        }

        if b.Set(Piece{}, s) == nil {
                t.Errorf("failed in set")
        }

        delete(b.board, s)
        if b.Set(Piece{}, s) != nil {
                t.Errorf("faield in set")
        }
}

func TestTake(t *testing.T) {
        b := NewBoard()
        s , _:= nthSquare(15)
        if _, err := b.Take(s); err == nil {
                t.Errorf("failed in take")
        }

        p := Piece{Type: King, Color : White}
        b.board[s] = p
        if res, err := b.Take(s); err != nil || res != p {
                t.Errorf("failed in take")
        }

        if _, err := b.Take(s); err == nil {
                t.Errorf("failed in take")
        }
}

func TestMove(t *testing.T) {
        b := NewBoard()
        from, _ := nthSquare(63)
        to, _ := nthSquare(22)
        b.board[to] = Piece{}
        b.board[from] = Piece{}
        err := b.Move(from, to)
        _, ok := err.(SquareOccupiedError)
        if err == nil || !ok  {
                t.Errorf("failed in move")
        }

        delete(b.board, to)
        fakeFrom, _ := nthSquare(53)
        err = b.Move(fakeFrom, to)
        _, ok = err.(NoPieceAtSquareError)
        if err == nil || !ok {
                t.Errorf("failed in move")
        }

        err = b.Move(from, to)
        if err != nil {
                t.Errorf("failed in move")
        }
}

func TestAt(t *testing.T) {
        b := NewBoard()
        s, _ := nthSquare(23)
        fake, err := nthSquare(64)
        if err != nil {
                t.Errorf("failed in nthSquare")
        }

        p := Piece{Type: King, Color: Black}
        b.board[s] = p
        if resP, ok := b.At(s); !ok || resP != p {
                t.Errorf("failed in At")
        }

        if _, ok := b.At(fake); ok {
                t.Errorf("failed in At")
        }

        delete(b.board, s)
        if _, ok := b.At(s); ok {
                t.Errorf("faield in At")
        }
}

func nthSquare(n uint) (Square, error) {
        s := FirstSquare()
        var err error
        for i := uint(0); i < n - 1; i++ {
                s, err = s.Next()
                if err != nil {
                        return Square{}, err
                }
        }

        return s, nil
}