package piece

import "testing"

func TestValidateWhite(t *testing.T) {
	c := White
	res := c.validate()
	if res != nil {
		t.Errorf("failed in validate white, err : %v", res.Error())
	}
}