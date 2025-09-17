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
