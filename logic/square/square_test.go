package square

import "testing"

const a  = File('a')
const h = File('h')

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

func TestFirstSquareFile(t *testing.T) {
	bs := FirstSquare()
        exp := a
        res := bs.file
	if exp != res {
		t.Errorf("exp : %v, res : %v", exp, res)
	}
}

func TestFirstSquareRank(t *testing.T) {
	bs := FirstSquare()
	if bs.rank != 1 {
		t.Errorf("failed to generate first square")
	}
}

func TestLastSquare(t *testing.T) {
	bs := LastSquare()
	if bs.file != h || bs.rank != Rank(8) {
		t.Errorf("failed to generated last square")
	}
}
